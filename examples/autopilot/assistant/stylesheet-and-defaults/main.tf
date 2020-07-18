resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "twilio-test"

  defaults = templatefile("${path.module}/defaults.json", {
    fallbackURLOrTask = "http://localhost/fallback"
  })
  stylesheet = templatefile("${path.module}/stylesheet.tpl.json", {
    voice = "Polly.Salli"
  })
}