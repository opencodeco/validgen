---
name: Application bug report
about: Create a bug report to help us fix an issue with Validgen
title: ''
labels: ''
assignees: ''

body:
  - type: markdown
    attributes:
      value: |
        Thank you for reporting an issue with ValidGen.
        Please try to fill in as much of the form below as possible. Fields marked with an asterisk (*) are required.

  - type: checkboxes
    id: read_getting_started
    attributes:
      label: Pre-check
      options:
        - label: I have read the [Contributing guide](../CONTRIBUTING.md)
          required: true

  - type: checkboxes
    id: check_duplicates
    attributes:
      label: Check for duplicates
      options:
        - label: I have searched for duplicate [issues](https://github.com/opencodeco/validgen/issues) (both open and closed).
          required: true

  - type: dropdown
    id: platform
    attributes:
      label: Your platform
      options:
        - Windows
        - Linux
        - MacOS
      default: 0
    validations:
      required: true

  - type: textarea
    id: expected
    attributes:
      label: Expected behaviour
      description: Describe clearly what you expected ValidGen to do.
      placeholder: E.g., ValidGen should generate code X or validate input Y correctly.
    validations:
      required: true

  - type: textarea
    id: actual
    attributes:
      label: Actual behaviour
      description: Describe clearly what actually happened in ValidGen.
      placeholder: E.g., ValidGen crashes or produces incorrect output.
    validations:
      required: true

  - type: textarea
    id: steps
    attributes:
      label: Steps to reproduce
      description: Provide step-by-step instructions to reproduce the bug.
      placeholder: What you did first, second, etc.
    validations:
      required: false

  - type: textarea
    id: additional_context
    attributes:
      label: Additional context
      description: Any extra information, logs, or screenshots that may help.
      placeholder: Add context here...
    validations:
      required: false
