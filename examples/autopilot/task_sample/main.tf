resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
}

resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = twilio_autopilot_task.task.assistant_sid
  task_sid      = twilio_autopilot_task.task.sid
  language      = "en-US"
  tagged_text   = "test"
}

data "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = twilio_autopilot_task_sample.task_sample.assistant_sid
  task_sid      = twilio_autopilot_task_sample.task_sample.task_sid
  sid           = twilio_autopilot_task_sample.task_sample.sid
}

data "twilio_autopilot_task_samples" "task_samples" {
  assistant_sid = twilio_autopilot_task_sample.task_sample.assistant_sid
  task_sid      = twilio_autopilot_task_sample.task_sample.task_sid
}
