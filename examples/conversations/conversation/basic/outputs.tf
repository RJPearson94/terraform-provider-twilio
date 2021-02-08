output "service" {
  description = "The Generated Conversations Service"
  value       = twilio_conversations_service.service
}

output "conversation" {
  description = "The Generated Conversations Conversation"
  value       = twilio_conversations_conversation.conversation
}

