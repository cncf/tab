# Reference Architecture Submission Process

The Cloud Native Reference Architecture site is a community‑driven project led by the CNCF End User Technical Advisory Board. It showcases real architectures and the stories behind them - how they were designed, the trade-offs made, and how they map to cloud‑native principles in practice.

## Why it’s worth doing

Beyond helping the broader community, submitting your architecture:
* Highlights the work your team is already doing in production.
* Provides concrete examples that ground the CNCF reference architecture and make it more useful for others.
* Gives you structured feedback from peers and TAB members before the architecture becomes “frozen” in people’s minds.

## Submission

Two paths for submission.

### Open Issue

If you're still thinking on how it might look like, start with an issue in our repo, this makes it easy to get started.

Here’s the flow:
* Open a new issue for a [Reference Architecture Submission](https://github.com/cncf/tab/issues/new?template=reference-architecture.yml)
* Use the template to capture:
  * A short context section: what this system does and the main requirements (scale, latency, availability, compliance, etc.).
  * A high‑level architecture view: the main building blocks, how they interact, and how they embody cloud‑native characteristics like being distributable, observable, portable, interoperable, and resilient.
  * The CNCF projects and other technologies you’re using, and why.
  * A few key lessons learned (things you’d want your past self or peers to know).

That’s all you need to start. It does not have to be perfect. Rough diagrams, partial descriptions, and “work in progress” are totally fine for the first pass.

We’ll iterate together on diagrams, wording, and structure to align with the overall reference architecture site and get it ready for publication on architecture.cncf.io.

### Pull Request

If you have all the info you need, check [our template](https://github.com/cncf/tab/blob/main/operations/templates/template-reference-architecture.md) for reference architectures.

Then branch from our repo and add yours [in here](https://github.com/cncf/architecture/tree/main/content/en/architectures). Submit the pull request and we'll pick it up from there.

