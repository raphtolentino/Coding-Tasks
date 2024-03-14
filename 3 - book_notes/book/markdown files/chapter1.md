
- [1. Section 2: The basics of Information Security](#1-section-2-the-basics-of-information-security)
  - [1.1. Information Security](#11-information-security)
    - [1.1.1. Malicious software](#111-malicious-software)
    - [1.1.2. Unauthorised access](#112-unauthorised-access)
    - [1.1.3. System Failure](#113-system-failure)
    - [1.1.4. Social Engineering](#114-social-engineering)
  - [1.2. Proactive Security Plan](#12-proactive-security-plan)
    - [1.2.1. Physical](#121-physical)
    - [1.2.2. Technical](#122-technical)
    - [1.2.3. Administrative](#123-administrative)
    - [1.2.4. Anti-Malware software](#124-anti-malware-software)
    - [1.2.5. Data Backup](#125-data-backup)
    - [1.2.6. Encryption](#126-encryption)
    - [1.2.7. Data Removal](#127-data-removal)
- [2. Section 3 : Think Like a Hacker](#2-section-3--think-like-a-hacker)
    - [2.0.1. White Hats](#201-white-hats)
    - [2.0.2. Black Hats](#202-black-hats)
    - [2.0.3. Grey Hats](#203-grey-hats)
    - [2.0.4. Blue Hats](#204-blue-hats)
- [3. Glossary](#3-glossary)
- [4. Chapter Review Activities](#4-chapter-review-activities)
    - [4.0.1. Review Questions (9/11)](#401-review-questions-911)
    - [4.0.2. Questions](#402-questions)
    - [4.0.3. Answers](#403-answers)

## 0.1. Sources

## 0.2. Chapter 1

### 0.2.1. Section 1 - Foundation Topics

1. CIA Triad

What is the CIA triad ?

- The purpose of the CIA triad is to provide relevant security knowledge using easy words.

### 0.2.2. **********************Confidentiality**********************

Characteristics of preventing the disclosure of information/data to unauthorised persons. For example Passport id number is confidential as its contain sensitive information.

Its important to follow regulations:

****HIPAA Regulations:****

- **PII (Personally Identifiable Information) r**efers to any data about an individual that could be used to identify them.
- ******************************************PHI (Personally Health Information)****************************************** refers to any health data about an individual that could be used to identify them.
- **********Classified or Sensitive Information********** is also a factor that involves trade secrets, research and other IPs.

### 0.2.3. ******************Integrity******************

Characteristics of data or system content being completed, correct and consistent.

- ********************Data Integrity******************** refers to the assurance that data has not been altered by an unauthorised manner. This requires that the data is free from modifications, errors, loss of information; and it's recorded, used and maintained safely that ensures completeness.
- ****************System Integrity**************** refers to the assurance that data is secure in a safe manner free from unauthorised access or manipulation, both intentional or not. It involves creating and maintaining good configuration.

The data or system integrity state can modify by comparing the baseline against the current state. If these two factors match, the system is secured.

### 0.2.4. ****************************Availability****************************

- Characteristics of timely and reliably access to information and being able to use it. For authorised users to have timely, reliable access to data and information.
- This does not mean that data/system need to be available 100% of the time. But, the data/system need to meet timely and reliability requirements.

## 0.3. The AAA of Computer Security

What is the AA of Computer Security:

### 0.3.1. Authentication

Characteristics of validating a person established identity with proof and confirmed by a system. It can be granted by:

- Username and Password
- Biometric data
- Tokens

There’s two authentication method types used in security.

- **Single Factor Authentication (SFA)** refers to single methods that does not require extra security to login-in to a system.
- Multi **Factor Authentication (MFA)** refers to single methods that require extra security to login-in to a system.

### 0.3.2. Authorisation

Characteristics of giving users access to certain data of a building. It is used after authentication and can be provided in several ways:

- Permission
- Access control Lists
- Time of Day restrictions
- Login or Physical restrictions

### 0.3.3. Accounting

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

## 0.4. Non-repudiation

A legal term defined as protection against a individual falsely denying having performed a particular action, such as:

- Creating information
- Approving Information
- Sending and receiving messages

## 0.5. Privacy

The Right to control the distribution of information about them self.

---

# 1. Section 2: The basics of Information Security

## 1.1. Information Security

What is Information Security:

It refers to the act of protecting data and information system from ****************************************************************************************************************unauthorised access, unlawful modification and disruption, corruption and destruction.****************************************************************************************************************

### 1.1.1. Malicious software

It refers to malware software intended for unlawful practises. It can be:

- Trojan Horses : Software hidden with bad code
- Spyware: Software with tracking code
- Rootkits : Software with hidden code to attack system files
- Adware: Software to fill the computer with advertisements
- Ransomware : Locks the computer and demands payment.
- Crypto-malware: mines crypto in the background

### 1.1.2. Unauthorised access

It having unauthorised access to computer resources without knowledge or permission of the owner. This includes:

- Approaching the system
- Trespassing
- Communicating
- Storing and Retrieving data
- Intercepting data
- Others

### 1.1.3. System Failure

### 1.1.4. Social Engineering

It refers to computer crashes or individual application failure. It can be caused by

- **user error**
- **hardware failure.**
- **malicious activity**

It refers to the act of manipulating users into revealing confidential information or to perform actions that is bad for the users.  

- It can be ******************phishing****************** emails
- Asking for ************************************************************personal information (or money).************************************************************
- Others

## 1.2. Proactive Security Plan

It is important to look after 3 categories of controls.

### 1.2.1. Physical

Physical Items like:

- Alarm systems
- CCTV cameras
- Locks
- ID cards
- Security Guards
- Others

### 1.2.2. Technical

Technical Items like:

- Smart Cards
- Access Control Lists
- Encryption
- Network Authentication

### 1.2.3. Administrative

Policies and Procedures designed to prevent this stuff to happens:

- Security awareness training
- Contingency plans
- Disasters recovery plans
- Others

### 1.2.4. Anti-Malware software

It refers to software aimed to protect a computer from different malware, by detecting and removing them. It can be:

- Antivirus
- Anti-Spyware software

### 1.2.5. Data Backup

To create a copy of the original data. Used to remove data after an attack or other compromise, or system failure.

It can be done with RAID 1,5,6 and 10 versions to prevent hardware failure:

- Windows Backup and Restore
- Bacula

### 1.2.6. Encryption

It refers to changing information of a file using a algorithm (cyphers) to make unreadable to anyone that does not have the correct key. The most popular encryption method is:

- AES (Encryption Standard)
- HTTP Secure (HTTPS)
- Secure/Multipurpose Internet Mail Extensions
- Pretty Good Privacy

### 1.2.7. Data Removal

The act of removing data. It goes far beyond of removing by deleting data. It involves purging the residue and dealing with ensuring that it cannot be recovered.

# 2. Section 3 : Think Like a Hacker

### 2.0.1. White Hats

Non-Malicious hackers that attempts to hack info a computer system with the owner permissions. Usually this type of hackers has a contractual agreement to enter the system.

A Ethical hacker is a expert at breaking into systems and can hack into systems on behalf of the system owner with prior permissions. They use PenTesting and Intrusion testing tools to target a network.

### 2.0.2. Black Hats

Malicious hackers, that attempts to break into computers without authorizations. They are the hacker group responsible for:

- Identity theft
- Piracy
- Credit card fraud
- Other Crimes

The penalty for these crimes are severe.

### 2.0.3. Grey Hats

Individuals that is between black and white hat hacking. They break into systems then notify the admin of the system that that they were successfully in doing do. They do not do anything else than breaking in.

Some offers to fix those issues, but they are considered **green hats**.

### 2.0.4. Blue Hats

Individuals that are asked to attempt to hack into a system by an organisation, but they are not employed by them.

# 3. Glossary

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

# 4. Chapter Review Activities

### 4.0.1. Review Questions (9/11)

### 4.0.2. Questions

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

### 4.0.3. Answers

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
