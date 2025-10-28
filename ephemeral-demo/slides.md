# Ship Faster and Test Smarter with Ephemeral Environments

## Slide 1: Problem

"Reviewing PRs without a live environment slows teams and hides bugs. I will show a preview environment that anyone can copy today."

Action: Show repo tree quickly.

---

## Slide 2: Architecture (20 seconds)

"Buildx and Bake produce a tagged image. Compose deploys the stack. Traefik routes by host. Teardown is one command."

[Diagram showing: Buildx + Bake → Tagged Image → Compose + Traefik → Live URL]

---

## Slide 3: Value and next steps (30 seconds)

Bullets to say out loud:

- Instant PR previews improve feedback and reduce risk
- Template is portable to a VPS or CI with a few env vars
- Add a nightly reaper and a PR workflow to automate lifecycle

---

## Slide 4: QR (10 seconds)

Ask the audience to scan and star.

[QR code linking to repository]
