name: "Reference Architecture Submission"
description: "Submit a Reference Architecture for Community Review"
title: "[Reference Architecture: ]"
labels: ["area/reference-architecture"]
body:
 - type: markdown
   attributes:
     value: |
       **Purpose of This Issue**
 
       This is a submission for a proposal of a reference architecture, as [defined here](https://github.com/cncf/tab/blob/main/process/reference-architectures.md). It should not contain the full content but a summary to evaluate its relevance to the community.
 - type: input
   id: name
   attributes: 
     label: Name
     description: A name for this architecture submission. Suggestions include the name of the end user or domain as well as any specific area of interest (CI/CD, Edge, etc)
     placeholder: Edge Deployments at Organization XYZ
   validations:
     required: true
 - type: input
   id: organization
   attributes:
     label: Organization
     description: Architectures must be submitted by at least one CNCF Member Company
     placeholder: List of Organizations
   validations:
     required: true
 - type: textarea
   id: teams
   attributes:
     label: Team(s)
     description: Describe the teams involved in this architecture
     placeholder: Developer Platforms, ...
   validations:
     required: true
 - type: textarea
   id: overview
   attributes:
     label: Overview and Goals
     description: A brief overview of the goals of the implementation
     placeholder: Edge deployment for sensor tracking
   validations:
     required: true
 - type: textarea
   id: projects
   attributes:
     label: Projects
     description: List of projects being used in the proposed architecture (include all projects, also non CNCF projects)
     placeholder: Kubernetes, Helm, ArgoCD, ...
   validations:
     required: true
 - type: textarea
   id: evolution
   attributes:
     label: Planned Evolution
     description: Any evolution planned for this reference architecture
     placeholder: Improvements in the area of networking and monitoring of edge devices
   validations:
     required: true
