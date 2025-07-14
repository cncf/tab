name: "Reference Architecture Submission"
description: "Submit a Reference Architecture for Community Review"
title: "[Reference: ]"
labels: ["reference-architecture"]
body:
 - type: markdown
   attributes:
     value: |
       **Purpose of This Issue**
 
       This is a submission to the reference architecture process.

       Add more information here for submitters...
 
 - type: input
   id: name
   attributes: 
     label: Name
     description: What is the name of this architecture submission?
     placeholder: e.g., My Reference Architecture
   validations:
     required: true
 - type: input
   id: affiliation
   attributes:
     label: Affiliation
     description: Architectures must be submitted by at least one CNCF Member Company
     placeholder: List of Companies
   validations:
     required: true
 - type: textarea
   id: description
   attributes:
     label: Organizational Description
     description: Describe your Organization 
     placeholder: A brief description of your organization. e.g. a paragraph from an about page on your company's website
   validations:
     required: true
 - type: textarea
   id: team
   attributes:
     label: Team Description
     description: Describe your Team
     placeholder: Do you work for a central IT organization in your company?  or for a team of a specific department?
   validations:
     required: true
 - type: textarea
   id: platform-topology-internal
   attributes:
     label: Internal Topology
     description: TBD
     placeholder: Describe your Internal Platform Engineering Team Topology
   validations:
     required: true
 - type: textarea
   id: platform-topology-external
   attributes:
     label: External Topology
     description: TBD
     placeholder: Describe your External Platform Engineering Team Topology
   validations:
     required: true
 - type: textarea
   id: overview
   attributes:
     label: Overview and Goals
     description: TBD
     placeholder: Brief overview of your architecture and any potential goals you are trying to achieve with it?
   validations:
     required: true
 - type: textarea
   id: services
   attributes:
     label: Projects and Services
     description: TBD
     placeholder: Can you expand on why you are using those projects/services?
   validations:
     required: true
 - type: textarea
   id: successes
   attributes:
     label: Successful Lessons Learned
     description: TBD
     placeholder: What has worked well?
   validations:
     required: true
 - type: textarea
   id: improvements
   attributes:
     label: Improvements that could be made  
     description: TBD
     placeholder: What has not worked well?
   validations:
     required: true
 - type: textarea
   id: glue
   attributes:
     label: Custom work and glue required
     description: TBD
     placeholder: What sort of "glue" have you had to develop to enable usage of your architecture? What have you done to have to get everything to work well together?e .g. have you written a bunch of helper scripts in bash? or maybe your own custom controller to manage a rollout?
   validations:
     required: true
 - type: textarea
   id: evolution
   attributes:
     label: Updates and Changes
     description: TBD
     placeholder: Has your architecture evolved? What lessons did you learn from previous iterations?
   validations:
     required: true
 - type: textarea
   id: future
   attributes:
     label: What's next?
     description: TBD
     placeholder: What's next for your architecture? What are you looking to do next?
   validations:
     required: true
