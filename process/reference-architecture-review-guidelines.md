# Reference Architecture Review Guidelines

This document provides comprehensive guidelines for reviewing CNCF reference architecture submissions. It serves as both a review checklist and an agent skill document for conducting independent, thorough evaluations.

## Agent Instructions

You are an assistant to a technical reviewer. Be consise and curt, do not use exaggarated language, state the facts clearly. A human reviewer will make the ultimate decision, focus on surfacing the data.

## Review Principles

**Independent Reviews:**
- Each reference architecture MUST be reviewed independently against these guidelines
- DO NOT compare architectures to each other during review
- Each review stands on its own merits
- Comparative analysis only appropriate when explicitly requested
- Testing guidelines means evaluating guideline effectiveness per architecture, not ranking architectures

**Review Process:**
- Weight all criteria equally unless specifically instructed otherwise
- Provide evidence-based assessment with specific examples
- Create GitHub issues in fork (castrojo/tab) for iterative review, always keep these issues up to date so that the issues reflect the current state of this review document.
- Issues should be updated automatically as reviews are refined, ensure only one issue per reference architecture exists. Each issue is a running copy.
- DO NOT create issues in upstream cncf/tab without explicit permission
- DO NOT TAG ANY INDIVIDUAL

### Review Philosophy

The TAB review process serves as a **data-surfacing tool** to help the Technical Advisory Board make informed decisions about reference architecture submissions. Reviews are:

- **Opt-in contributions** - Organizations voluntarily share their architectures to benefit the community
- **Collaborative assessments** - Identifying development opportunities, not passing judgment
- **Maturity-based** - Evaluating readiness across multiple dimensions
- **Context-aware** - Considering organizational goals and industry representation

The goal is to provide TAB with comprehensive information about submission readiness, not to categorically approve or reject contributions.

**Rating Scale:**

We use a **maturity model** approach to assess each criterion:

- **🟢 Established**: Demonstrates strong maturity in this dimension
- **🟡 Developing**: Shows progress with opportunities for enhancement
- **⚪ Emerging**: Early stage with significant development recommended

This positive framing recognizes that submissions are at different stages of their journey and focuses on constructive pathways forward.

**Overall Assessment:**

After evaluating each criterion, provide a **narrative assessment** that:

- Summarizes maturity across all dimensions
- Identifies key strengths and development opportunities
- Describes readiness for TAB consideration
- Suggests constructive pathways forward

**Example:** "This submission demonstrates established technical depth (🟢) and strong CNCF relevance (🟢), with developing metrics (🟡) and emerging organizational context (⚪). The architecture would benefit from additional outcome data and industry positioning before TAB review."

---

## 1. Initial Submission Checks

### 1.1 Content Quality

**🟢 Established:**
- Issue template filled out completely and correctly
- Submitted to correct repository with proper labels
- All required fields have substantive responses (not placeholders)
- Grammar and text suitable for professional publication
- Not raw LLM output (shows human editing and domain expertise)

**🟡 Developing:**
- Most fields completed but some sections need expansion
- Shows promise with opportunities for refinement
- Some sections feel template-driven but show original content

**⚪ Emerging:**
- Missing required fields or placeholder content
- Submitted to wrong location or incorrect format
- Would benefit from refinement in grammar and clarity
- Content would benefit from substantial development

**Historical Context:**

Published reference architectures typically have:
- 3,000-5,000 words of substantive technical content
- Multiple architecture diagrams (3-5 minimum)
- Specific metrics and scale data
- Clear problem statement and solution description

Example from Adobe: 4,200 words, 5 architecture diagrams, specific metrics (360+ clusters, 22K+ apps), clear platform engineering narrative.

---

## 2. CNCF End User Member Status

### Submitter Engagement

**🟢 Established:**
- Submitter is a CNCF End User Member in good standing
- Listed in cncf/landscape repository as End User Member
- Demonstrated engagement with CNCF (events, case studies, community participation)
- Clear path to increased engagement if not already deeply involved

**🟡 Developing:**
- Organization is exploring CNCF membership
- Some engagement but not formally a member yet
- Strong potential for future engagement

**⚪ Emerging:**
- No CNCF connection or membership
- No clear path to membership or engagement
- May benefit from end user perspective emphasis

**Why This Matters:**

Reference architectures should come from end users who are part of the CNCF community. This ensures:
- Authentic real-world implementation experience
- Alignment with CNCF values and open source principles
- Credibility and trust within the community
- Opportunity for ongoing engagement and knowledge sharing throughout the organization and its members

**Verification Steps:**

1. Check CNCF member list: https://www.cncf.io/about/members/
2. Verify end user status in cncf/landscape repository:

```bash
# Verify end user membership
curl -s https://raw.githubusercontent.com/cncf/landscape/master/landscape.yml | \
  grep -A 5 "name: \"Organization Name\"" | \
  grep "enduser: true"
```

If the command returns output, the organization is confirmed as an end user member.

3. Look for past KubeCon presentations or case studies
4. Check GitHub contributions to CNCF projects (if claimed)

---

## 3. CNCF Relevance and Audience Value

### CNCF Project Usage

**🟢 Established:**
- Clear identification of which CNCF projects are being used (minimum 3-4 core projects)
- Specific versions and configurations documented
- Integration patterns and architecture decisions explained
- Usage goes beyond basic Kubernetes (shows advanced adoption)

**🟡 Developing:**
- CNCF projects identified but details lacking
- Heavy reliance on proprietary or vendor-specific solutions
- Basic Kubernetes usage without ecosystem integration

**⚪ Emerging:**
- CNCF projects not clearly identified
- Primarily vendor/proprietary solution with minimal open source
- No clear connection to CNCF technology stack

**Community Engagement Test: "Would People Attend This KubeCon Talk?"**

Ask yourself:
- Is this topic interesting to the broader CNCF community?
- Does it showcase innovative use of CNCF projects?
- Would attendees learn something actionable?
- Is it "hot" tech OR critical infrastructure work that deserves recognition?

Both cutting-edge innovations AND solid production infrastructure stories are valuable.

**CNCF Project Collaboration:**

Strong submissions demonstrate:
- Active contribution to CNCF projects they use
- Engagement with project maintainers and communities
- Bug reports, feature requests, or code contributions
- Participation in SIGs or working groups
- Sharing lessons learned with project communities via [KubeCon talks](https://www.youtube.com/@cncf) 

**Historical Context:**

From 179 published case studies:
- **83%** mention Kubernetes extensively (not just "we use k8s")
- **32%** document Prometheus/observability integration
- **28%** describe Helm usage for packaging
- **18%** implement service mesh (Istio, Linkerd)

Strong submissions show ecosystem adoption beyond just Kubernetes.

**Technology Alignment:**

Reference to CNCF End User Technology Radar:
- Core components should primarily use "Adopt" ring technologies
- "Trial" technologies acceptable with clear rationale and production validation
- "Assess" or unlisted technologies require strong justification

See `reference-architectures-data/cncf-enduser-radar.md` for current radar status.

---

## 4. Problem Statement and Solution Value

### Problem Definition

**🟢 Established:**
- Clear problem statement with business context
- Scale and complexity of problem well explained
- Quantified pain points (e.g., "deployments took 4 hours", "MTTR was 6 hours")
- Clear explanation of why existing solutions were insufficient

**🟡 Developing:**
- Problem stated but lacks specificity
- Business impact not quantified
- Scale indicators vague or missing

**⚪ Emerging:**
- No clear problem statement
- Generic "needed to modernize" without specifics
- Solution presented without explaining the problem

**Solution Uniqueness:**

Strong submissions demonstrate:
- **Unique approach**: Novel solution or innovative adaptation of existing patterns
- **Lessons learned**: Insights valuable to others facing similar challenges
- **Best practices**: Reusable patterns and anti-patterns identified
- **Industry relevance**: Applicable to specific vertical or broadly across industries

**Problem Categories:**

Reference architectures typically address:
- **Production challenges**: Scaling, reliability, performance
- **Developer experience**: Productivity, onboarding, workflow efficiency  
- **Operational efficiency**: Cost reduction, automation, incident response
- **New capabilities**: Enabling new products or business models
- **Compliance/security**: Meeting regulatory or security requirements

**Historical Examples:**

**Adobe** (Platform Engineering):
- **Problem**: 250+ teams, inconsistent tooling, slow onboarding (weeks)
- **Scale**: 360+ clusters, 22K+ apps, multiple clouds
- **Solution**: Unified platform with self-service capabilities
- **Outcome**: Onboarding reduced to days, consistent developer experience

**Typical Financial Services** (from case studies):
- **Problem**: Regulatory compliance, data isolation, security requirements
- **Scale**: Multi-region, millions of transactions/day
- **Solution**: Zero-trust architecture, policy-as-code, comprehensive audit trails
- **Outcome**: Achieved compliance certifications, reduced audit time 70%

---

## 5. Scale and Adoption Indicators

### Scale Metrics

**🟢 Established:**
- Specific quantitative metrics provided across multiple dimensions
- Clear indication of production deployment (not prototype)
- Scale appropriate to demonstrate pattern validity
- Growth trajectory and adoption timeline documented

**🟡 Developing:**
- Some metrics provided but incomplete picture
- Scale unclear or appears inflated
- Production status ambiguous

**⚪ Emerging:**
- No scale metrics or only vague statements ("many services", "large scale")
- Appears to be pilot or proof-of-concept, not production
- Numbers don't align with rest of submission

**Key Dimensions:**

**Technical Scale:**
- Cluster/node counts
- Service/application counts  
- Container/pod counts
- Data volumes processed
- Transaction/request volumes
- Geographic distribution (regions, data centers)
- Impact on the business, metrics used to measure this

**People Scale:**
- Number of teams using the platform
- Total engineers involved
- Platform team size
- Organizational reach (business units, subsidiaries)
- Impact on the business, metrics used to measure this

**Operational Scale:**
- Deployment frequency
- Number of deployments per day/week
- Services onboarded over time
- Growth rate and trajectory
- Impact on the business, metrics used to measure this
  
**Scale Benchmarks from Published Case Studies:**

**Large-Scale (Enterprise):**
- 100+ Kubernetes clusters
- 1,000+ services/applications
- 100+ development teams
- Multi-cloud, multi-region
- Examples: Adobe (360+ clusters), Spotify (1,200+ services), Apple (100K+ nodes)

**Mid-Scale (Growth Company):**
- 10-50 Kubernetes clusters
- 100-500 services
- 10-50 development teams
- Single or multi-cloud
- Examples: New York Times (40+ clusters, 300+ apps)

**Emerging Scale (Startup/SMB):**
- 5-20 Kubernetes clusters
- 20-100 services
- 5-20 teams
- Single cloud, possibly multi-region
- Examples: UK Government Digital Service (30+ services, 100+ teams)

**Note:** Scale alone doesn't determine value. An emerging-scale architecture with unique insights can be as valuable as a large-scale one.

---

## 6. Organizational and Industry Representation

### Stakeholder Analysis

Reference architectures should clearly identify and address multiple stakeholder perspectives:

**Primary Stakeholders:**

**Platform Team:**
- What problems does this architecture solve for platform engineers?
- How does it reduce operational burden?
- What automation and tooling is provided?
- How is the platform maintained and evolved?
- Impact on the business, metrics used to measure this

**Application Developers:**
- How does this improve developer experience?
- What self-service capabilities are provided?
- How is onboarding simplified?
- What guardrails and golden paths exist?

**Security Team:**
- How are security requirements addressed?
- What compliance frameworks are satisfied?
- How is policy enforcement automated?
- What audit and monitoring capabilities exist?

**Business Leadership:**
- What business outcomes are achieved?
- How is ROI demonstrated?
- How are the business metrics measured?
- What competitive advantages are gained?
- How does this enable business strategy?

**SRE/Operations:**
- How is reliability ensured?
- What observability is built-in?
- How are incidents detected and resolved?
- What disaster recovery capabilities exist?

**🟢 Established:**
- Multiple stakeholder perspectives addressed
- Clear value proposition for each stakeholder group
- Trade-offs and decision rationale explained
- Organizational change management considerations included

**🟡 Developing:**
- Some stakeholders addressed but others missing
- Value propositions stated but not well documented
- Limited discussion of organizational impact

**⚪ Emerging:**
- Only technical perspective (no business or organizational context)
- Single stakeholder view (e.g., only developers or only ops)
- No consideration of change management or adoption challenges

### Industry Representation

**Current Industry Distribution (from 179 case studies):**

- **Well-Represented (10+ case studies):**
  - Financial Services: 26 (14.5%)
  - Software/Technology: 26 (14.5%)
  - Telecommunications: 18 (10.1%)
  - Retail/E-commerce: 16 (8.9%)
  - Media/Entertainment: 12 (6.7%)

- **Moderately-Represented (5-10 case studies):**
  - Transportation/Automotive: 10 (5.6%)
  - Government/Public Sector: 8 (4.5%)

- **Under-Represented (<5 case studies):**
  - Healthcare/Life Sciences: 4 (2.2%)
  - Manufacturing: 3 (1.7%)
  - Energy/Utilities: 3 (1.7%)
  - Education: 2 (1.1%)

**Review Consideration:**

- 🟢 Submissions from under-represented industries receive favorable consideration
- 🟢 Healthcare, Manufacturing, Energy, Education particularly valuable
- 🟡 Submissions from well-represented industries must demonstrate unique value
- ⚪ Generic submissions from over-represented industries may not add community value

**Geographic Diversity:**

Current distribution:
- North America: ~40%
- Europe: ~32%
- Asia-Pacific: ~23%
- Rest of World: ~5%

Submissions from under-represented geographies (Latin America, Africa, Middle East) are particularly valuable.

### Organizational Context

**🟢 Established:**
- Industry, size, and regulatory context clearly explained
- Specific requirements that influenced architecture decisions
- Team structure and organizational model described
- Cultural and change management aspects addressed

**🟡 Developing:**
- Basic organizational info provided but lacks depth
- Context doesn't clearly connect to architecture decisions
- Missing key organizational aspects (size, teams, structure)

**⚪ Emerging:**
- No organizational context provided
- Generic statements not specific to organization
- Architecture presented in vacuum without business context

---

## 7. Quality Standards and Documentation

### Architecture Diagrams and Visual Representations

**🟢 Established:**
- Multiple diagrams showing different aspects (3-5 minimum)
- High-level architecture overview
- Network topology and data flow
- Security boundaries and controls
- CI/CD pipeline and deployment flow
- Diagrams are clear, professional, and properly labeled
- Legend provided for symbols and notation
- Suitable for "typical CTO" to understand quickly

**🟡 Developing:**
- Some diagrams provided but coverage incomplete
- Diagrams present but could be clearer or more detailed
- Missing key architectural views (e.g., security, networking)

**⚪ Emerging:**
- No diagrams or only one high-level diagram
- Diagrams would benefit from refinement in clarity and labeling
- Too many diagrams creating confusion rather than clarity
- Diagrams don't match text description

**Documentation Quality:**

**🟢 Established:**
- Logical flow from problem → solution → implementation → outcomes
- Technical depth appropriate for expert practitioners
- Specific examples and code snippets where helpful
- Clear explanation of "why" not just "what"
- Architectural Decision Records (ADRs) or similar documentation referenced

**🟡 Developing:**
- Content present but organization could be improved
- Some sections too shallow, others too detailed
- Flow could be more logical

**⚪ Emerging:**
- Organization would benefit from improvement
- Would benefit from additional technical depth
- Missing critical technical information
- Inconsistent level of detail across sections

### Reputation and Professional Standards

**"Does this demonstrate technical excellence as a team?"**

Strong reference architectures:
- Showcase technical excellence and thoughtful design
- Demonstrate deep expertise and lessons learned
- Reflect positively on the organization and team
- Would make leadership proud to see in headlines
- Would work well as KubeCon presentation by appealing to the greater cloud native community

**Community Discussion Test:**

Ask yourself:
- Would this generate discussion among end users?
- Are there insights that would spark conversation?
- Would specific industry segments find this particularly relevant?
- Does this advance the state of practice?

**Professional Quality:**

Content should:
- Be suitable for professional publication
- Show careful editing and review
- Avoid marketing speak or vendor promotion
- Maintain technical accuracy
- Use industry-standard terminology correctly

---

## 8. Metrics and Measurable Outcomes

### Methodology and Measurement

**🟢 Established:**
- Specific, quantitative metrics provided (not vague statements)
- Before/after comparisons with concrete numbers
- Clear methodology for how metrics were measured
- Multiple outcome categories (technical, business, operational)
- Metrics credible and aligned with stated scale

**🟡 Developing:**
- Some metrics provided but methodology unclear
- Mix of quantitative and vague qualitative statements
- Metrics seem aspirational rather than actual
- Limited before/after comparison

**⚪ Emerging:**
- No metrics or only qualitative statements ("faster", "better")
- Methodology completely unclear
- Metrics don't align with submission scale or seem inflated
- No baseline for comparison

**Metric Categories:**

**Technical Metrics:**
- Deployment frequency (e.g., "from weekly to multiple daily")
- Lead time for changes (e.g., "from 3 weeks to 2 days")
- Mean time to recovery (e.g., "from 4 hours to 15 minutes")
- System reliability/uptime (e.g., "99.97% availability achieved")
- Performance improvements (e.g., "P95 latency reduced from 2s to 200ms")

**Business Metrics:**
- Time to market (e.g., "feature delivery 50% faster")
- Cost efficiency (e.g., "40% infrastructure cost reduction")
- Revenue impact (e.g., "enabled launch of 50 new products")
- Customer satisfaction (e.g., "NPS increased 20 points")

**Developer Productivity Metrics:**
- Onboarding time (e.g., "from 1 month to 1 week")
- Build/test time (e.g., "CI pipeline from 45 min to 8 min")
- Environment provisioning (e.g., "from 2 days to 10 minutes")
- Developer satisfaction scores

**Operational Metrics:**
- Resource utilization (e.g., "increased from 45% to 85%")
- Incident frequency (e.g., "reduced by 60%")
- Manual intervention reduction (e.g., "90% less manual configuration")

**Historical Examples:**

**Adobe** (Platform Engineering):
- "Reduced onboarding from weeks to days"
- "360+ clusters, 22,000+ applications"
- "250,000+ containers in production"
- "Unified developer experience across multiple clouds"

**Typical Financial Services** (from case studies):
- "10,000+ deployments per day" (Spotify scale)
- "MTTR reduced from 4 hours to 15 minutes"
- "99.99% uptime SLA achieved"
- "70% reduction in compliance audit time"

**What Makes Metrics Credible:**

- **Specific numbers** not ranges or approximations
- **Time-bound** (when were these achieved?)
- **Context provided** (how measured, what changed)
- **Reasonable** (aligned with stated scale and industry)
- **Multiple categories** (not just one metric type)

---

## Review Deliverable Format

### GitHub Issue Structure

When creating a review issue in `castrojo/tab` fork:

**Title Format:**
```
Review: [Architecture Name] - [Overall Assessment]
```

Example: `Review: Adobe Platform Engineering - STRONG PASS (4.5/5.0)`

**Issue Body Template:**

```markdown
# Reference Architecture Review: [Name]

**Reviewer:** [Your name/handle]
**Review Date:** [YYYY-MM-DD]
**Submission Link:** [Link to original submission issue]

## Overall Assessment

**Rating:** [STRONG PASS / PASS / PARTIAL PASS / NEEDS WORK]
**Score:** [X/8 criteria at ✅ Pass]

**Summary:** [2-3 paragraph summary of overall findings]

---

## Detailed Criteria Evaluation

### 1. Initial Submission Checks
**Rating:** [✅ / ⚠️ / ❌]

[Evidence-based assessment with specific examples]

### 2. CNCF End User Member Status
**Rating:** [✅ / ⚠️ / ❌]

[Evidence-based assessment with specific examples]

### 3. CNCF Relevance and Audience Value
**Rating:** [✅ / ⚠️ / ❌]

[Evidence-based assessment with specific examples]

### 4. Problem Statement and Solution Value
**Rating:** [✅ / ⚠️ / ❌]

[Evidence-based assessment with specific examples]

### 5. Scale and Adoption Indicators
**Rating:** [✅ / ⚠️ / ❌]

[Evidence-based assessment with specific examples]

### 6. Organizational and Industry Representation
**Rating:** [✅ / ⚠️ / ❌]

[Evidence-based assessment with specific examples]

### 7. Quality Standards and Documentation
**Rating:** [✅ / ⚠️ / ❌]

[Evidence-based assessment with specific examples]

### 8. Metrics and Measurable Outcomes
**Rating:** [✅ / ⚠️ / ❌]

[Evidence-based assessment with specific examples]

---

## Recommendations

**Strengths:**
- [Bullet list of key strengths]

**Areas for Improvement:**
- [Bullet list of specific improvements needed]

**Questions for Submitter:**
- [Questions requiring clarification or additional information]

---

## Next Steps

[Recommended next steps based on overall assessment]

- **STRONG PASS**: Recommend acceptance and publication
- **PASS**: Minor revisions suggested, then accept
- **PARTIAL PASS**: Specific revisions required before acceptance
- **NEEDS WORK**: Substantial revision needed, resubmit after improvements

---

## References

- [Link to CNCF case studies research]
- [Link to Technology Radar data]
- [Link to submission template]
- [Any other relevant references]
```

### Review Best Practices

**Evidence-Based Assessment:**
- Quote specific sections from submission
- Reference line numbers or section headers
- Provide examples from historical data to support judgments
- Avoid subjective statements without evidence

**Constructive Feedback:**
- Be specific about what needs improvement
- Suggest concrete actions or examples
- Acknowledge strengths as well as weaknesses
- Maintain professional, respectful tone

**Actionable Recommendations:**
- Provide clear next steps
- Prioritize critical issues vs. nice-to-haves
- Suggest specific resources or examples to reference
- Make recommendations feasible and realistic

---

## Agent-Specific Instructions

### For AI Review Agents

When conducting a reference architecture review:

1. **Read the submission thoroughly**
   - Review the GitHub issue with all fields
   - Read any linked documentation or diagrams
   - Check for referenced external resources

2. **Research the submitter**
   - Verify CNCF membership status
   - Check cncf/landscape repository for end user member status
   - Look for past KubeCon presentations or case studies
   - Review any GitHub contributions if claimed

3. **Consult historical data**
   - Reference `reference-architectures-data/cncf-case-studies-research.md`
   - Reference `reference-architectures-data/cncf-enduser-radar.md`
   - Compare to published reference architectures at https://architecture.cncf.io/

4. **Evaluate each criterion independently**
   - Use the maturity model rating scale (🟢 / 🟡 / ⚪) consistently
   - Provide specific evidence for each rating
   - Quote from submission to support assessments
   - Reference historical examples for context

5. **Provide narrative overall assessment**
   - Summarize maturity across all dimensions
   - Identify key strengths and development opportunities
   - Describe readiness for TAB consideration
   - Suggest constructive pathways forward

6. **Create comprehensive GitHub issue**
   - Use the review deliverable template
   - Include all sections with substantive content
   - Provide actionable recommendations
   - Link to relevant resources

7. **Maintain independence**
   - Review each submission on its own merits
   - Do NOT compare to other submissions unless explicitly requested
   - Apply criteria consistently across all reviews

8. **Update issue as needed**
   - If new information is provided, update the review
   - Track changes in review as submission evolves
   - Maintain conversation with submitter if clarifications needed

9. **Clarify confusing acronyms while recognizing industry partnerships**
   - If source documents mention "CNOE" (Cloud Native Operational Excellence), always spell it out on first reference to avoid confusion with "CNCF"
   - Example: "Published blog on Cloud Native Operational Excellence (CNOE).io" or "Active in CNOE (Cloud Native Operational Excellence) cohort"
   - Recognize legitimate industry partnerships and affiliations (Linux Foundation projects, standards bodies, industry consortiums) when relevant to the architecture story
   - CNCF commonly partners with other organizations - these relationships add value and context
   - Focus reviews primarily on CNCF project usage, CNCF community engagement, and CNCF End User membership, while acknowledging broader ecosystem participation

### Common Pitfalls to Avoid

❌ **Don't:**
- Rate based on organization size or fame
- Penalize for using non-"Adopt" technologies without considering rationale
- Compare submissions to each other
- Focus only on technical aspects (remember business and organizational context)
- Accept vague or generic statements without evidence
- Ignore industry representation considerations
- Use confusing acronyms like "CNOE" without spelling them out (easily confused with "CNCF")

✅ **Do:**
- Apply criteria consistently and fairly
- Provide constructive, actionable feedback
- Acknowledge unique value even at different scales
- Consider submitter's specific context and constraints
- Value both innovation and solid production stories
- Recognize industry under-representation as valuable
- Acknowledge legitimate industry partnerships and ecosystem participation (Linux Foundation, standards bodies, consortiums)

---

## Related Documents

- **Process Workflow**: `process/reference-architectures.md` - 6-step review and publication process
- **Submission Template**: `.github/ISSUE_TEMPLATE/reference-architecture.yml` - What submitters fill out
- **Case Studies Research**: `process/reference-architectures-data/cncf-case-studies-research.md` - Historical data
- **Technology Radar**: `process/reference-architectures-data/cncf-enduser-radar.md` - Technology adoption patterns
- **Agent Instructions**: `AGENTS.md` - General agent guidance and repository context

---

**Document Version:** 2.0
**Last Updated:** February 2026
**Maintainer:** CNCF TAB Reference Architecture Review Team
