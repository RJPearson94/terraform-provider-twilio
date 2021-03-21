resource "twilio_video_composition_hook" "composition_hook" {
  friendly_name = "Test Composition Hook"
  audio_sources = ["*"]
  format        = "mp4"
}
