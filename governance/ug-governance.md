# CNCF User Group Governance

**NOTE:** This document is a starting point. It is expected to change and be refined as User Groups take

This document outlines the scope and responsibilities of the CNCF User Groups

The CNCF User Groups are comprised of four types of groups that serve different purposes or have different scopes:

**User Group (UG):** Primary organizational unit that is aligned with an industry vertical, problem domain, or a group that serves a common need of the CNCF End User community.

**Initiative:** A lightweight organizational unit used for UG or TAB work that is tightly scoped and time-bound. Initiatives are typically run and staffed from within the UG, with the TAB governing the initiative proposals and ensuring correct alignment and duration.

**Technical Community Group:** A lightweight topic or domain-focused group used to serve as a rallying point for community members to discuss, share knowledge, and coordinate potential future initiatives.

- [CNCF User Group Governance](#cncf-user-group-governance)
- [Keywords](#keywords)
- [User Groups (UGs)](#user-groups-\(ugs\))
  - [Goals](#goals)
  - [Formation](#formation)
  - [Requirements](#requirements)
  - [Roles, Terms, and Elections](#roles,-terms,-and-elections)
    - [Chair](#chair)
      - [Requirements](#requirements-1)
      - [Duties](#duties)
      - [Activity Expectations](#activity-expectations)
      - [Escalations](#escalations)
    - [Other Roles](#other-roles)
- [Initiatives](#initiatives)
  - [Goals](#goals-1)
  - [Formation](#formation-1)
    - [UG Initiatives](#ug-initiatives)
    - [TAB Initiatives](#tab-initiatives)
  - [Requirements](#requirements-2)
  - [Roles](#roles)
- [Technical Community Groups](#technical-community-groups)
  - [Formation](#formation-2)
  - [Roles, Terms, and Elections](#roles,-terms,-and-elections-1)
    - [Organizer](#organizer)
      - [Requirements](#requirements-3)
      - [Duties](#duties-1)
      - [Activity Expectations](#activity-expectations-1)
      - [Escalations](#escalations-1)
    - [Other Roles](#other-roles-1)

## Keywords

This document uses [RFC2119 keywords] to indicate requirement levels as clearly
and concisely as possible. RFC2119 has long been used by many governance and
standards bodies, such as the [Internet Engineering Task Force (IETF)] and the
[Kubernetes project].

- **MUST** - This word, or the terms "**REQUIRED**" or "**SHALL**," means that
  the item is required.
- **MUST NOT** - This phrase, or the phrase "**SHALL NOT**", means that the item
  is not permitted or must be avoided.
- **SHOULD** -  This word, or the adjective "**RECOMMENDED**," means the item is
  suggested but may be ignored for valid reasons.
- **SHOULD NOT** - This phrase, or the phrase "**NOT RECOMMENDED**," means the
  item is acceptable under certain circumstances.
- **MAY** - This word, or the adjective "**OPTIONAL**," means the item is truly
  optional.


[RFC2119 keywords]: https://www.ietf.org/rfc/rfc2119.txt
[Internet Engineering Task Force (IETF)]: https://en.wikipedia.org/wiki/Internet_Engineering_Task_Force
[Kubernetes project]: https://github.com/kubernetes/community/blob/master/committee-steering/governance/sig-governance-requirements.md


## User Groups (UGs) 

A CNCF UG will oversee and coordinate the interests pertaining to a logical area of needs across projects, other CNCF groups, and users. They are long-lived groups that report directly to the TAB. Areas covered by UGs are focused on industry problem domains or support the needs of other UGs and the TAB itself.


### Goals

* Represent and serve the needs of the user community in a particular domain  
* Identify gaps in the CNCF project portfolio. Find and attract projects to fill these gaps.  
* Educate and inform users with unbiased, effective, and practically useful information scoped to cloud native.  
* Solicit and propose publication of community sourced domain specific reference architecture(s), best practices, how-to’s and other artifacts produced by the community for the community  
* Provide feedback to CNCF projects on usability, reliability, and performance

### Formation

UGs are formed when there is a documented, outstanding gap in technical domain coverage, needs, and direction within cloud-native. UGs are expected to be very long lived groups that perform standing functions for projects and their adopters.

Due to the longevity of a UG, existing groups interested in becoming a UG need to engage the TAB. The TAB will review the request against the current needs and movement of the ecosystem and make a determination.

Example: A domain-focused community group has regularly submitted and completed initiatives that cover gaps in the current UG structure. After a discussion with the TAB, a charter with defined scopes is drafted. The TAB then votes on the UG charter PR, and the UG is formed once the vote passes.

### Requirements

* **MUST** have a charter that outlines scope, potential overlaps with other UGs, and any operational deviations from UG Governance.  
* **MUST** enumerate any roles within the UG and the responsibilities of each along with any role membership requirements and lifecycle.  
* **MUST** maintain UG “metadata” \- a list of leads, and initiatives.  
* If applicable, **MUST** document services provided to other UGs or projects.  
* **MUST** hold at least one regularly scheduled public monthly meeting, recorded and uploaded to YouTube.  
* **MUST** provide periodic updates to the TAB and TAB liaisons on their health and initiatives.  
* **SHOULD** provide regular ecosystem updates on the domain, such as through the regular CNCF community meeting, KubeCon presentations, or to the TOC directly.  
* **SHOULD** organize KubeCon \+ CloudNativeCon presence; such as CFPs, Project Pavilion booths etc.  
* **SHOULD** form and document recurring ongoing work or services.  
* **SHOULD** produce well researched and informed technical white papers, best practices, and other ecosystem supportive assets.  
* **SHOULD** surface gaps or ecosystem trends.

### Roles, Terms, and Elections

#### Chair

A Chair is an individual who is responsible for both providing subject matter expertise of a UG’s defined domain as well as overseeing and driving the planning, execution, and accountability of efforts within the UG. UG chairs are expected to share responsibilities equally. They should consider rotating/taking turns performing chair duties to ensure consistency, availability, and continuity over the course of Chair terms and availability.

Chairs build unity in the purpose of the group for which they govern and oversee. This includes sufficient consideration of ideas, topics, and issues presented, facilitation of healthy discussion, delegation of tasks (as appropriate), and ensuring their group's adherence to both the [technical leadership principles] and the TAB’s operating principles.

Chairs ensure agreed tasks are carried out and that group decisions have consensus. They are responsible for reporting progress to their parent body (TAB).

Chairs also primarily perform administrative functions including collecting and compiling topics for the agenda, chairing the meeting, ensuring that quality meeting notes or minutes are published, and follow-up actions tracked and resolved.

The expected minimum time commitment is at least 2-4 hours of work per week but expected to increase depending on involvement in initiatives, during times with a high volume of requests, and/or KubeCon preparations and activities

##### Requirements

* **MUST** be approved by the TAB.  
* **MUST** adhere to the [Technical Leadership Principles].  
* Membership **MUST** be tracked alongside the TAG metadata.  
* **MUST** have prior experience within the CNCF, but **MAY** be approved by the TAB for an explicitly documented reason such as extensive domain expertise, formation of a new governing body (e.g., new UG, or demonstrating significant aptitude for people & community management.

##### Duties

* **SHOULD** drive charter changes (including creation) to get community buy-in but **MAY** delegate content creation to other leads or community members.  
* **SHOULD** define how priorities and commitments are managed and set a cap on concurrent efforts of the group per quarter. **MAY** delegate to other leads as needed.  
* **MUST** review technical outputs or artifacts from the group for accuracy. This includes white papers, code, other forms of technical documentation, etc.  
* **MUST** ensure leadership roles are documented, accurate, and updated in any applicable locations (GitHub teams, mailing lists, etc).  
* **MUST** coordinate communication and be a connector with other community groups like UGs and the TAB, but **MAY** delegate the actual communication and creation of content to other contributors where appropriate.  
* **SHOULD** facilitate meetings but **MAY** delegate to other Leads or active community members.  
* **SHOULD** mentor future Leads as part of active and continuous succession planning.  
* **MUST** regularly triage their issues & PRs but **MAY** delegate to other contributors to fulfill these responsibilities.  
* **MUST** organize presence at one or more KubeCon \+ CloudNativeCon (CFPs, Project Pavilion booth, etc) and approve content presented by or about the UG, but **MAY** delegate to other contributors.  
* **MUST** regularly provide status updates on their groups and their initiatives.  
* **MUST** coordinate service desk requests on behalf of the UG.  
* **MUST** keep the UG liaisons informed of the group’s activities and progress against the charter.

##### Activity Expectations

* Chairs MUST be present or ensure representation in the monthly public TAB meeting reporting on recent activities when relevant  
* Chairs serve 2 year terms, and **MAY** submit for re-election.  
* Chairs **MAY** decide to step down at any time and propose a replacement to complete their term.  
  * The candidate **SHOULD** be supported by a majority of the current active leads, with support backed up by documentation, such as links to GitHub activity or documented deliverables.  
  * The candidate is **REQUIRED** to be approved by the TAB.  
* Chairs taking an extended leave of 1 or more months **MUST** coordinate with other leads to ensure the role is adequately covered during their leave.  
* Chairs going on leave for 1-3 months **MAY** work with other Leads to identify a temporary replacement.  
* Active Leads are **REQUIRED** to notify the TAB and receive confirmation of notice before pursuing removal of an inactive or Lead found to be in dereliction of duty.  
* Inactive Leads **MUST** be removed through a super-majority vote of the TAB, following an internal group discussion and escalation.

##### Escalations

* Lead membership disagreements **MAY** be escalated to the TAB or CoCC as applicable.

#### Other Roles

UGs **MAY** have other roles, but they are **REQUIRED** to be documented in the charter and require the approval of their parent governing body.

All additional roles are **REQUIRED** to have the following minimum requirements:

* **MUST** be documented with a description, scope, and responsibilities.  
* **MUST** have a documented lifecycle policy including how they are appointed, any potential terms, and an inactivity policy.  
* **MUST** be tracked alongside the other group metadata.  
* **MUST** adhere to the [Technical Leadership Principles].  
* **SHOULD** remain active and responsive in their Roles.

## Initiatives

Initiatives are short-term, time bound work that have a pre-defined objective and exit criteria. Some examples include white papers, one-off reports or working with the TAB on a specific initiative (e.g. Feedback Loops). They do not have a charter, or defined leads, but can have separate meetings if it is beneficial to the initiative completing its goal. Each initiative **MUST** have a UG Chair or TAB member assigned to oversee and ensure timely completion and closure of the work within the scope of the parent governing group's charter.

### Goals

* Serve as a lightweight organizational unit to coordinate work.  
* Recognize community members who have helped execute and deliver on the initiative.  
* Ensure high-velocity completion of scoped effort with clear outcomes aligned with the CNCF, TAB, and the particular UGs charter.

### Formation

There are two types of initiatives with different formation requirements: UG initiatives and TAB initiatives.

**A Note on Collaboration with TAGs and TOC**

Initiatives that are more project focused are encouraged to be created under the TOC or one of the TAGs. They are the best equipped to provide guidance and shepherd project focused work. Similarly, the TAGs and TOC are encouraged to create an initiative in cooperation with the UGs or TAB when an initiative is more user-focused e.g. cluster operator best practices.

#### UG Initiatives

General initiatives can be any "*piece of work*," such as white papers, reports, etc., that align with the scope of a UG. The Chairs manage UG Initiatives and **MUST** have at least one assigned as the sponsor to oversee the initiative's execution. They **MAY** delegate some of the responsibility of organizing and shepherding the initiative, but are still ultimately responsible for it.

UG initiatives can span multiple UGs when something spans multiple tech domains, however one is ultimately assigned as the responsible group for execution and reporting. Initiatives that span many UGs, or have no clear owner should be proposed as TAB initiatives.

#### TAB Initiatives

The TAB may create and own initiatives when the scope of work spans the majority of UGs or coordinate work around specific goals. In these instances, at least one TAB member **MUST** be assigned to organize, shepherd, and report on the initiative's status.

### Requirements

* **MUST** be documented, including goals, milestones, deliverables, and any exit criteria (e.g., a tracking GitHub issue).  
* Initiatives **MUST** have more than two participants not including the assigned lead for oversight.  
* **SHOULD** actively be worked on. An initiative that has made little progress in 3 months should be canceled or paused until other contributors can be identified.  
* Initiatives **SHOULD** not last longer than 2 quarters but **MAY** request an extension from their parent governing body with reasonable justification.  
* **MUST** check in at least monthly to assess the initiative's continued need, progress, and ability to meet.  
* **SHOULD** acknowledge everyone who actively helped complete the initiative.  
* 

### Roles

Assigned Chairs or TAB members serve as the general oversight of initiatives depending on the parent governing body.

## Technical Community Groups

Technical Community Groups (TCGs) are a part of the broader [CNCF Community Groups Program]. They are a lightweight topic or domain-focused group used as a rallying point for community members to discuss, share knowledge, and coordinate potential future initiatives. They are free to propose and shepherd initiatives when there is an explicit deliverable or objective they wish to work on.

Community Groups have [much broader exposure], and have a suite of tools tailored for community management, making them the best place for discussion-oriented groups.

Additionally, it is likely that future UGs will likely start as community groups.

A community group applying to become a UG:

* **MUST** have completed numerous initiatives and cover a gap in the currently defined groups  
* **MUST** be approved by the TAB  
* **MUST** have its leadership established through the same processes outlined in the respective section above for UGs


[CNCF Community Groups Program]: https://github.com/cncf/communitygroups
[much broader exposure]: https://community.cncf.io/


### Formation

TCG's have minimal governance requirements and follow the formation process outlined in [Community Groups README], with the additional step of requiring TAB approval.

**NOTE:** If a potential TCG could be created under the purview of either the TOC or TAB; the TCG organizers should reach out to the chairs of the TOC and TAB to discuss which group would be the most appropriate.


[Community Groups README]: https://github.com/cncf/communitygroups?tab=readme-ov-file#how-to-apply


### Roles, Terms, and Elections

#### Organizer

An Organizer is an individual domain expert who is responsible for the overall administration and execution of the TCG, such as scheduling and hosting meetings, promoting the group, and coordinating the creation of initiatives.

##### Requirements

**MUST** adhere to the [Technical Leadership Principles].  
**MUST** adhere to the [Community Group Organizer Requirements].  
**SHOULD** remain active and responsive in their role.

##### Duties

* **SHOULD** facilitate meetings but **MAY** delegate to other Leads or active community members.  
* **SHOULD** mentor future Leads as part of active and continuous succession planning.  
* **MUST** coordinate communication and be a connector with other community groups like TAGs and the TOC, but MAY delegate the actual communication and creation of content to other contributors where appropriate.

##### Activity Expectations

* Organizers serve 2-year terms.  
* Organizers **MUST** remain active and responsive.  
* Organizers **MAY** decide to step down at any time and propose a replacement to complete their term.  
* The candidate **SHOULD** be supported by a majority of the current active organizers.  
* Organizers taking an extended leave of 1 or more months **MUST** coordinate with other Organizers to ensure the role is adequately covered during their leave.  
* Organizers going on leave for 1-3 months **MUST** work with the other Organizers to identify a temporary replacement.  
* Organizers who are inactive, unreachable, or on leave for more than 3 months without a designated temporary replacement MUST be considered in dereliction of duty.  
* Active Organizers are **REQUIRED** to notify the community group organizers at community-groups@cncf.io and receive confirmation of notice before pursuing removal of an inactive or Organizer found to be in dereliction of duty.

##### Escalations

* Organizer disagreements MAY be escalated to the TAB, CoCC or CNCF Staff facilitating the Community Groups program as applicable.

#### Other Roles

TCG's MAY have other roles, but they are **REQUIRED** to be documented in an easily discovered area, such as the group's page on Bevy; the platform used by Community Groups.

* **MUST** be documented with a description, scope, and responsibilities.  
* **MUST** adhere to the Technical Leadership Principles.  
* **SHOULD** remain active and responsive in their roles.


<!-- Common Links -->

[Technical Leadership Principles]: https://github.com/cncf/toc/blob/main/PRINCIPLES.md#technical-leadership-principles
[Community Group Organizer Requirements]: https://github.com/cncf/communitygroups/blob/main/organizers.md