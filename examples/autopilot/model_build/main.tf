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

resource "twilio_autopilot_model_build" "model_build" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid

  triggers = {
    redeployment = sha1(join(",", tolist([
      twilio_autopilot_task_sample.task_sample.sid,
      twilio_autopilot_task_sample.task_sample.language,
      twilio_autopilot_task_sample.task_sample.tagged_text,
    ])))
  }

  lifecycle {
    create_before_destroy = true
  }

  polling {
    enabled = true
  }
}
