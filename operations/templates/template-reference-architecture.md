---
title: Template Reference Architecture
date: 2026-03-28

org_name: ACME
org_team: ACME Platform Team
org_url: https://acme.myorg
org_logo_filename: images/logo.svg
contact: Jane Doe, John Doe
email: jane@acme.org, john@acme.org

org_description: |
  ACME is a fictional organization used as a placeholder for the CNCF TAB reference architecture template.

org_size: "10,000+"
user_size: "10,000+"

industries:
- Research
- AddHere
tags: 
- cloud_native
- kubernetes
- addhere

reference_architectures:
- research
- ai
---

## Relevant Projects

{{< cardpane >}}

  {{< card header="ArgoCD" >}}
  [![argo logo](https://github.com/cncf/artwork/raw/main/projects/argo/horizontal/color/argo-horizontal-color.svg)](https://www.cncf.io/projects/argo/)
  - **Using since:** 2024
  - **Current version:** v3.2.0

  ArgoCD is used to manage deployments of all services across multiple clusters and environments.
  {{< /card >}}

  {{< card header="Kubernetes" >}}
  [![kubernetes logo](https://raw.githubusercontent.com/cncf/artwork/main/projects/kubernetes/icon/color/kubernetes-icon-color.svg)](https://www.cncf.io/projects/kubernetes/)
  - **Using since:** 2024
  - **Current version:** 1.31.x

  Kubernetes provides the required workload scheduling and orchestration for the diverse workloads running in the platform.
  {{< /card >}}

  {{< card header="PROJECTXYZ" >}}
  [![xyz logo](https://raw.githubusercontent.com/cncf/artwork/main/projects/xyz/icon/color/xyz-icon-color.svg)](https://www.cncf.io/projects/xyz/)
  - **Using since:** 
  - **Current version:** 

  Copy / paste this card to add other projects you rely on.
  {{< /card >}}

{{< /cardpane >}}

## TLDR; Synopsis

( replace with your own )

This reference architecture describes a fictional deployment for ACME, supported by its internal platform team. It has revolutionized the "ways of working" in the organization, allowing the migration from traditional VM deployments to a cloud native infrastructure.

In particular, this architecture targets:
* A declarative definition of workloads and services, delegating to the orchestrator the role of placing the required replicas in a highly available manner
* ( other things you should highlight for your own reference architecture )

## Organization

( replace with your own )

ACME is a fictional organization supporting the template for submission of reference architectures in the CNCF TAB. ACME has no real internal structure, but if it did, it would be made of an infrastructure team focusing on the low level datacenter deployment, a platforms team offering higher level tools for teams to deploy their services in a highly scalable manner, and a devex team that ensures a consolidated development experience across all departments.

ACME builds and offers nothing, but if it did, it would be awesome.

## Teams

( replace with your own )

**Infrastructure**, maintains our robust datacenter deployments, focusing on procurement but also hardware lifecycle and repairs. It offers the interfaces needed for our cluster deployments.

**Platform**, the core of our Kubernetes solutions, ...

**Add your own teams**, highlighting the revelant ones for this architecture.

## Architecture

### Goals

( replace with your own )

This architecture targets different goals, including:
**Full GitOps and Automation**, covering the compute workloads but also the management of storage, network and DNS registrations, and all the resources that require lifecycle management.

**Support for heterogenous resources**, allowing the deployment over CPUs, GPUs and including resources both on-premises and external.

**Add your own items**, focusing on the key goals for your architecture.

### Architecture Overview

( replace with your own )

The diagram below shows an overview of our architecture.

**Infrastructure layer** focusing on the management and lifecycle of hardware, with Metal3 as the interface for resource provisioning of compute.

**Advanced scheduling layer**, with Kubernetes as the main orchestrator but with a custom scheduler plugin capable of making decisions based on internal knowledge. The plugin talks to an existing backend at ACME that can provide the required optimizations.

**Add and replace with your own**, highlighting any particular configurations and optimizations that improved the performance of the deployment, and all the design principles that were followed to end up with this reference architecture. If needed, split this into multiple sub-sections (i.e. compute, storage, etc).

## Can you expand on why you are using those projects/services?

( replace as appropriate... this section is optional )

For each of the selected projects, consider expanding on what made you choose them for that purpose and include any additional relevant information that led to that choice - including any comparison you might have made with other projects.

## What works particularly well

( replace with your own )

The **integration of GPUs** was particularly easy thanks to the advanced support and driver availability integrating with DRA, including the scheduling, allocation and monitoring of the resources.

**Add your own items**, mentioned all areas where the architecture works particularly well or the integration was easy.

## What needs improvement

( replace with your own )

**Low latency networking** is a key requirement but our current architecture relies on hostNetwork, dropping the network isolation between workloads. Moving to DRA also for the networking part is one of the goals to be accomplished in the coming months.

**Add your own items**, highlighting things where the results are not fully satisfactory, pain points you would like to point out or gaps in the ecosystem that led to additional work.

## What sort of "glue" have you had to develop?

( replace with your own )

**Tenant specific configurations** could not be easily exposed in the resource definitions, leading the a large number of mutating policies that help enforcing them. The logic is kept in a dedicated helm chart.

**Add your own items**, highlighting any special component or scripts you had to build yourself or around the existing projects to achieve your goals - things like special operators, integration with internal storage, network or identity, etc.

## How did the Architecture Evolve

( replace with your own, optional )

## What's next for your architecture?

**Expanding to external resources** while still offering a seamless experience to our users, now that they're used to the cloud native tools and interfaces. Projects like XYZ should allow us to achieve it in the coming months, and a few prototypes are already in place.

**Add your own items** covering your roadmap and how you plan to expand your architecture.

## Key Takeaways / Lessons

( replace with your own )

**Fully declarative setup**: applying GitOps best practices for all areas of our deployment is now possible, thanks to the easy integration of CRDs for internal workload definitions.

**Add your own items**: highlight the key take-aways for others to learn from your experience.

## Discussion

End user members may participate in the [discussion thread](https://github.com/cncf/enduser-private/discussions/??) for this architecture.

