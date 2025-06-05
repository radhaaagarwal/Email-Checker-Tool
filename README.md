# Email-Checker-Tool

**Domain security reconnaissance built with Go**

A lightweight DNS intelligence tool that reveals the email security posture of any domain. Uncover MX records, SPF policies, and DMARC configurations to assess email infrastructure and detect potential vulnerabilities.

---

## Why This Tool Matters

Email spoofing costs businesses billions annually. This tool cuts through the noise by analyzing the three critical security mechanisms that protect email domains. Whether you're auditing corporate infrastructure or investigating suspicious domains, you get the intelligence you need.

## Core Features

- **MX Record Discovery** - Identify mail server infrastructure
- **SPF Policy Analysis** - Extract anti-spoofing configurations  
- **DMARC Reconnaissance** - Reveal authentication policies
- **Batch Processing** - Analyze hundreds of domains instantly
- **CSV Export** - Clean output for reporting and analysis

---

## Real-World Usage

### Corporate Security Audit
```bash
# Check your company's email infrastructure
echo -e "company.com\nmail.company.com\nsupport.company.com" | go run main.go
```

### Threat Intelligence
```bash
# Analyze suspicious domains
cat suspicious_domains.txt | go run main.go > threat_analysis.csv
```

### Deliverability Assessment
```bash
# Verify marketing domain configuration
echo "newsletter-domain.com" | go run main.go
```

---

## Output Intelligence

The tool generates structured CSV intelligence:

```csv
domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord
google.com,true,true,v=spf1 include:_spf.google.com ~all,true,v=DMARC1; p=reject
insecure-site.com,true,false,,false,
well-configured.org,true,true,v=spf1 include:_spf.provider.com ~all,true,v=DMARC1; p=quarantine
```

**What Each Field Reveals:**
- `hasMX` - Can the domain receive emails?
- `hasSPF` - Anti-spoofing protection enabled?
- `spfRecord` - Complete sender authorization policy
- `hasDMARC` - Advanced authentication active?
- `dmarcRecord` - Email policy enforcement level

---

## Understanding Email Security

### MX Records
Mail server locations. No MX = no incoming email capability.

### SPF (Sender Policy Framework)
Prevents email spoofing by defining authorized sending servers.
```
v=spf1 include:_spf.google.com ~all  # Google Workspace authorization
```

### DMARC (Domain-based Message Authentication)
Advanced policy enforcement with three levels:
- `p=none` - Monitor only
- `p=quarantine` - Flag suspicious emails  
- `p=reject` - Block unauthorized emails

---

## Technical Architecture

Clean Go implementation leveraging standard libraries:

```go
scanner := bufio.NewScanner(os.Stdin)     // Efficient input processing
net.LookupMX(domain)                      // Mail server discovery
net.LookupTXT(domain)                     // SPF policy extraction
net.LookupTXT("_dmarc." + domain)         // DMARC configuration
```

**Key Strengths:**
- Zero external dependencies
- Memory-efficient streaming processing
- Concurrent DNS lookups for speed
- Non-fatal error handling

---


**Potential Enhancements:**
- DKIM record validation
- Concurrent processing optimization  
- JSON output format
- Historical change tracking
- Threat intelligence integration

---

