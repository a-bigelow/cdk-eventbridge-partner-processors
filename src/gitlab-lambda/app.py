"""
GitLab webhook receiver for EventBridge integration.
Validates GitLab webhook signatures and forwards events to EventBridge.
"""
import json
import os
import hmac
import hashlib
import boto3
from typing import Dict, Any

# Initialize AWS clients
secrets_client = boto3.client('secretsmanager')
events_client = boto3.client('events')

# Cache the secret to avoid repeated API calls
_cached_secret = None


def get_webhook_secret() -> str:
    """Retrieve the webhook secret from Secrets Manager."""
    global _cached_secret
    
    if _cached_secret is not None:
        return _cached_secret
    
    secret_arn = os.environ.get('GITLAB_WEBHOOK_SECRET_ARN')
    if not secret_arn:
        raise ValueError('GITLAB_WEBHOOK_SECRET_ARN environment variable not set')
    
    response = secrets_client.get_secret_value(SecretId=secret_arn)
    _cached_secret = response['SecretString']
    return _cached_secret


def verify_signature(payload: str, signature: str, secret: str) -> bool:
    """
    Verify GitLab webhook signature.
    GitLab uses X-Gitlab-Token header for webhook authentication.
    """
    if not signature:
        return False
    
    # GitLab uses a simple token comparison (not HMAC)
    return hmac.compare_digest(signature, secret)


def lambda_handler(event: Dict[str, Any], context: Any) -> Dict[str, Any]:
    """
    AWS Lambda handler for GitLab webhooks.
    
    Args:
        event: Lambda event containing the webhook request
        context: Lambda context
        
    Returns:
        API Gateway response
    """
    try:
        # Extract request details
        headers = event.get('headers', {})
        body = event.get('body', '')
        
        # Get GitLab webhook token from headers
        gitlab_token = headers.get('x-gitlab-token', headers.get('X-Gitlab-Token', ''))
        
        # Verify signature
        secret = get_webhook_secret()
        if not verify_signature(body, gitlab_token, secret):
            return {
                'statusCode': 401,
                'body': json.dumps({'message': 'Invalid signature'})
            }
        
        # Parse the webhook payload
        try:
            payload = json.loads(body) if body else {}
        except json.JSONDecodeError:
            return {
                'statusCode': 400,
                'body': json.dumps({'message': 'Invalid JSON payload'})
            }
        
        # Extract event type
        event_type = headers.get('x-gitlab-event', headers.get('X-Gitlab-Event', 'unknown'))
        
        # Send to EventBridge
        event_bus_name = os.environ.get('EVENT_BUS_NAME')
        if not event_bus_name:
            raise ValueError('EVENT_BUS_NAME environment variable not set')
        
        events_client.put_events(
            Entries=[
                {
                    'Source': 'gitlab.webhook',
                    'DetailType': event_type,
                    'Detail': json.dumps(payload),
                    'EventBusName': event_bus_name
                }
            ]
        )
        
        return {
            'statusCode': 200,
            'body': json.dumps({'message': 'Event processed successfully'})
        }
        
    except Exception as e:
        print(f'Error processing webhook: {str(e)}')
        return {
            'statusCode': 500,
            'body': json.dumps({'message': 'Internal server error'})
        }
