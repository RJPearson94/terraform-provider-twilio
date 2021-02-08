output "service" {
  description = "The Generated Conversations Service"
  value       = twilio_conversations_service.service
}

output "conversation" {
  description = "The Generated Conversations Conversation"
  value       = twilio_conversations_conversation.conversation
}

output "webhook" {
  description = "The Generated Conversations Webhook"
  value       = twilio_conversations_conversation_webhook.webhook
}

