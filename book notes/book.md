# Security Intro Notes [ISC2 and Book Notes]

- [Security Intro Notes \[ISC2 and Book Notes\]](#security-intro-notes-isc2-and-book-notes)
  - [Sources](#sources)
  - [Chapter 1](#chapter-1)
    - [Section 1 - Foundation Topics](#section-1---foundation-topics)
    - [**********************Confidentiality**********************](#confidentiality)
    - [******************Integrity******************](#integrity)
    - [****************************Availability****************************](#availability)
  - [The AAA of Computer Security](#the-aaa-of-computer-security)
    - [Authentication](#authentication)
    - [Authorisation](#authorisation)
    - [Accounting](#accounting)
  - [Non-repudiation](#non-repudiation)
  - [Privacy](#privacy)
- [Section 2: The basics of Information Security](#section-2-the-basics-of-information-security)
  - [Information Security](#information-security)
    - [Malicious software](#malicious-software)
    - [Unauthorised access](#unauthorised-access)
    - [System Failure](#system-failure)
    - [Social Engineering](#social-engineering)
  - [Proactive Security Plan](#proactive-security-plan)
    - [Physical](#physical)
    - [Technical](#technical)
    - [Administrative](#administrative)
    - [Anti-Malware software](#anti-malware-software)
    - [Data Backup](#data-backup)
    - [Encryption](#encryption)
    - [Data Removal](#data-removal)
- [Section 3 : Think Like a Hacker](#section-3--think-like-a-hacker)
    - [White Hats](#white-hats)
    - [Black Hats](#black-hats)
    - [Grey Hats](#grey-hats)
    - [Blue Hats](#blue-hats)
- [Glossary](#glossary)
- [Chapter Review Activities](#chapter-review-activities)
    - [Review Questions (9/11)](#review-questions-911)
    - [Questions](#questions)
    - [Answers](#answers)

## Sources

Book : [Title](../../Downloads/books/cyberbook.pdf)

## Chapter 1

### Section 1 - Foundation Topics

1. CIA Triad

What is the CIA triad ?

- The purpose of the CIA triad is to provide relevant security knowledge using easy words.

### **********************Confidentiality**********************

Characteristics of preventing the disclosure of information/data to unauthorised persons. For example Passport id number is confidential as its contain sensitive information.

Its important to follow regulations:

****HIPAA Regulations:****

- **PII (Personally Identifiable Information) r**efers to any data about an individual that could be used to identify them.
- ******************************************PHI (Personally Health Information)****************************************** refers to any health data about an individual that could be used to identify them.
- **********Classified or Sensitive Information********** is also a factor that involves trade secrets, research and other IPs.

### ******************Integrity******************

Characteristics of data or system content being completed, correct and consistent.

- ********************Data Integrity******************** refers to the assurance that data has not been altered by an unauthorised manner. This requires that the data is free from modifications, errors, loss of information; and it's recorded, used and maintained safely that ensures completeness.
- ****************System Integrity**************** refers to the assurance that data is secure in a safe manner free from unauthorised access or manipulation, both intentional or not. It involves creating and maintaining good configuration.

The data or system integrity state can modify by comparing the baseline against the current state. If these two factors match, the system is secured.

### ****************************Availability****************************

- Characteristics of timely and reliably access to information and being able to use it. For authorised users to have timely, reliable access to data and information.
- This does not mean that data/system need to be available 100% of the time. But, the data/system need to meet timely and reliability requirements.

## The AAA of Computer Security

What is the AA of Computer Security:

### Authentication

Characteristics of validating a person established identity with proof and confirmed by a system. It can be granted by:

- Username and Password
- Biometric data
- Tokens

There’s two authentication method types used in security.

- **Single Factor Authentication (SFA)** refers to single methods that does not require extra security to login-in to a system.
- Multi **Factor Authentication (MFA)** refers to single methods that require extra security to login-in to a system.

### Authorisation

Characteristics of giving users access to certain data of a building. It is used after authentication and can be provided in several ways:

- Permission
- Access control Lists
- Time of Day restrictions
- Login or Physical restrictions

### Accounting

Characteristics of tracking data, computer usage and network resources. It involves having proof if someone performed unauthorised actions. It involves:

- Logging
- Auditing
- Monitoring of data and resources

<aside>
🔑 It can be broken into 5 categories:

Something the users knows:

- Username and Password / Pin

Something the user has:

- Smart card and other security item

Something the user is:

- Biometric reading or Fingerprints or retina scans

Something the user does:

- Voice recognition or written signature

Somewhere the user is:

- GPS location.

</aside>

---

## Non-repudiation

A legal term defined as protection against a individual falsely denying having performed a particular action, such as:

- Creating information
- Approving Information
- Sending and receiving messages

## Privacy

The Right to control the distribution of information about them self.

---

# Section 2: The basics of Information Security

## Information Security

What is Information Security:

It refers to the act of protecting data and information system from ****************************************************************************************************************unauthorised access, unlawful modification and disruption, corruption and destruction.****************************************************************************************************************

### Malicious software

It refers to malware software intended for unlawful practises. It can be:

- Trojan Horses : Software hidden with bad code
- Spyware: Software with tracking code
- Rootkits : Software with hidden code to attack system files
- Adware: Software to fill the computer with advertisements
- Ransomware : Locks the computer and demands payment.
- Crypto-malware: mines crypto in the background

### Unauthorised access

It having unauthorised access to computer resources without knowledge or permission of the owner. This includes:

- Approaching the system
- Trespassing
- Communicating
- Storing and Retrieving data
- Intercepting data
- Others

### System Failure

### Social Engineering

It refers to computer crashes or individual application failure. It can be caused by

- **user error**
- **hardware failure.**
- **malicious activity**

It refers to the act of manipulating users into revealing confidential information or to perform actions that is bad for the users.  

- It can be ******************phishing****************** emails
- Asking for ************************************************************personal information (or money).************************************************************
- Others

## Proactive Security Plan

It is important to look after 3 categories of controls.

### Physical

Physical Items like:

- Alarm systems
- CCTV cameras
- Locks
- ID cards
- Security Guards
- Others

### Technical

Technical Items like:

- Smart Cards
- Access Control Lists
- Encryption
- Network Authentication

### Administrative

Policies and Procedures designed to prevent this stuff to happens:

- Security awareness training
- Contingency plans
- Disasters recovery plans
- Others

### Anti-Malware software

It refers to software aimed to protect a computer from different malware, by detecting and removing them. It can be:

- Antivirus
- Anti-Spyware software

### Data Backup

To create a copy of the original data. Used to remove data after an attack or other compromise, or system failure.

It can be done with RAID 1,5,6 and 10 versions to prevent hardware failure:

- Windows Backup and Restore
- Bacula

### Encryption

It refers to changing information of a file using a algorithm (cyphers) to make unreadable to anyone that does not have the correct key. The most popular encryption method is:

- AES (Encryption Standard)
- HTTP Secure (HTTPS)
- Secure/Multipurpose Internet Mail Extensions
- Pretty Good Privacy

### Data Removal

The act of removing data. It goes far beyond of removing by deleting data. It involves purging the residue and dealing with ensuring that it cannot be recovered.

# Section 3 : Think Like a Hacker

### White Hats

Non-Malicious hackers that attempts to hack info a computer system with the owner permissions. Usually this type of hackers has a contractual agreement to enter the system.

A Ethical hacker is a expert at breaking into systems and can hack into systems on behalf of the system owner with prior permissions. They use PenTesting and Intrusion testing tools to target a network.

### Black Hats

Malicious hackers, that attempts to break into computers without authorizations. They are the hacker group responsible for:

- Identity theft
- Piracy
- Credit card fraud
- Other Crimes

The penalty for these crimes are severe.

### Grey Hats

Individuals that is between black and white hat hacking. They break into systems then notify the admin of the system that that they were successfully in doing do. They do not do anything else than breaking in.

Some offers to fix those issues, but they are considered **green hats**.

### Blue Hats

Individuals that are asked to attempt to hack into a system by an organisation, but they are not employed by them.

# Glossary

<aside>
🔑 **State** : the condition of an entity at a specific point in time

B**aseline** : the lowest security level configuration allowed in a standard.

**Critically**:  Degree measure that an organisation depends on the information or its system for the success of a mission or a business function.

S**********ensitivity :********** the measure of importance by its owner; for the purpose of assigning correct protections.

</aside>

<aside>
🔑 **Script Kiddie**:  Individual with little technical knowledge that reuses code and script what is open of the internet.

**Hacktivist**: A attacker who has an agenda that might be good or bad.

**Organised Crime**:  Criminal groups that is motivated by money, using computers systems and hacking to gain information to access information and secrets.

**Advance Persistent Threat (APT)**: A set of computing attacking processes that target private organisations or nation states. It can be groups of governments that targets a specific entity.

</aside>

---

# Chapter Review Activities

### Review Questions (9/11)

### Questions

1. In information security, what are the three main goals? (Select the three best
answers.)
A. Auditing
B. Integrity
C. Non-repudiation
D. **Confidentiality**
E. Risk Assessment
F. Availability
2. To protect against malicious attacks, what should you think like?
A. Hacker
B. Network admin
C. Spoofer
D. Auditor
3. Tom sends out many e-mails containing secure information to other companies. What
concept should be implemented to prove that Tom did indeed send the e-mails?
A. Authenticity
B. Non-repudiation
C. Confidentiality
D. Integrity
4. Which of the following does the A in CIA stand for when it comes to IT security?
(Select the best answer.)
A. Accountability
B. Assessment
C. Availability
D. Auditing
5. Which of the following is the greatest risk when it comes to removable storage?
A. Integrity of data
B. Availability of data
C. Confidentiality of data
D. Accountability of data
6. When it comes to information security, what is the I in CIA?
A. Insurrection
B. Information
C. Indigestion
D. Integrity
7. You are developing a security plan for your organization. Which of the following is
an example of a physical control?
A. Password
B. DRP
C. ID card
D. Encryption
8. A user receives an e-mail but the e-mail client software says that the digital
signature is invalid and the sender of the e-mail cannot be verified. The would-be
recipient is concerned about which of the following concepts?
A. Confidentiality
B. Integrity
C. Remediation
D. Availability
9. Cloud environments often reuse the same physical hardware (such as hard drives)
for multiple customers. These hard drives are used and reused when customer
virtual machines are created and deleted over time. What security concern does this
bring up implications for?
A. Availability of virtual machines
B. Integrity of data
C. Data confidentiality
D. Hardware integrity
10. Which of the following individuals uses code with little knowledge of how it
works?
A. Hacktivist
B. Script kiddie
C. APT
D. Insider
11. When is a system completely secure?
A. When it is updated
B. When it is assessed for vulnerabilities
C. When all anomalies have been removed
D. Never

### Answers

1. B, D, F
2. A
3. B
4. C
5. C
6. D
7. C
8. B
9. C
10. B
11. D

---
