
---
name: BLIP-3
about: Introduce versioning to state objects across workflows and worksteps
title: "[BLIP-3] Trustmesh"
labels: 'State object versioning, Trustmesh, Baseledger'
assignees: ''

---

title: **BLIP-3 Trustmesh** - Versioning for state objects across workflows and worksteps

description: Real-life use-cases of baselining could require a way to version the state objects which are subject to baselining, so that a single state object with its historical changes could be tracked among multiple workflows\worksteps. Introducing this functionality would enable a standardized way to version the state objects as well as track their connection to subsequent state objects from the same domain, 
thus giving the baseline participant a clear overview of the history and version changes of a single domain transaction. This concept was originaly developed for Baseledger.

author: martenjung, stefanschmidt, steffankostic, ognjenkurtic

---

# Abstract

Current specification of the standard does not include option to version state objects across multiple workflows and\or worksteps although there are many business uses for this functionality. We propose to extend the BPI transaction with optional fields that would allow for this kind of versioning and tracking of a single state object as it is being baselined. These fields would be used to create a trustmesh entry, an entity that stores the baseline 'iteration' of a single state object, with reference to a previous 'iteration' of baselining. Trustmesh entry would be stored in the BPI Storage coomponent and made avaiable through the BPI API, enabling the system of record to visualize history of all iterations - a  Trustmesh.

# Motivation

The motivation behind Baseledger is to offer a blockchain solution with a certain degree of performance, service quality and compliance. Therefore, Baseledger aims to be positioned as Baseline-L2 to Ethereum. 
We now propose the concept of versioning and a synchronized versioning history that exists in Baseledger - Trustmesh - to be incorporated in Baseline.

The sequence of using Baseledger from the view of business participants looks as follows:

1.  Alice and Bob enter a workgroup    
2.  Alice wants to send a BusinessObject (BO) to Bob, this is a SyncronizationRequest    
3.  Alice converts the BO into a "SyncTree" (The leafs are the flattened JSON Tokens, the nodes are the calculated hashes of 2 underlying nodes)
4.  Alice packs the root proof and some metadata about the BO into a “BaseledgerPayload” of Type “Suggestion”, encrypts it and stores it in the blockchain    
5.  Alice sends an Offchain-Message to Bob, including the complete SyncTree and a reference to the BlockchainTransaction    
6.  Bob receives that message, looks up the referenced Blockchain Transaction, and decrypts the payload (he can do that because he is part of the same workgroup as the sender Alice)    
7.  Bob compares the proof stored in the blockchain with the one of the synctree and recalculates the complete SyncTree    
8.  If hashes match, Bob sends the Untokenized BO JSON to his system of record    
9.  Bob’s system of record decides about the Feedback (Approval or Rejection)    
10.  Bob creates BaseledgerPayload of type “Feedback”, referencing the original transaction and holding a positive or negative Feedback value, encrypts that payload and stores it in the blockchain    
11.  Bob sends an Offchain-Message to Alice, including the rootproof of the original Synctree and a reference to the blockchain transaction    
12.  Alice receives and unpacks the Feedback and updates her system of record to decide about the next step (e.g. a new version of a rejected document, or the next step in the workflow)
    
This complete history of requests, feedbacks, (new) versions and different worksteps are stored as TrustMesh entries. The TrustMesh defines the complete relation of different worksteps, their versions and approvals of one workflow. TrustMesh holds the references to BusinessObjects in the system of recrod, to TransactionIDs in the blockchain and to feedback gathered from Business Participants.

Currently, this versioning of state objects is not covered by the Baseline standard, and the only way a system of record can have a trace of baselining of a single business object through multiple workflows is to keep a local store of message and transaction ids etc. 

This approach leaves much to individual implementations, meaning various differences in the means of expectations and potential requirements to the standard. Having a standard for a common feature often needed in business communication reduces this future noise and complexity. 

# Specification

One proposal is to extend the BPI transaction to include the following optional fields:

**ReferencedTransactionId** (UID) - reference to a previous transaction dealing with the same state object
**StateObjectId** (UID) - a unique identifier of the business object in the originator's system of record

After the execution of a workstep and in case these fields are present, the output of the execution would contain a trustmesh entry - an entry that is then encrypted and stored in the Processing Layer storage.

Trustmesh entry would contain the following information:

**TransactionId** (UID) - unique identifier of the transaction
**ReferencedTransactionId** (UID) - identifier of the transaction with the same state object preceding this transaction
**StateObjectId** (UID) - unique identifier of the state object in the originator system of record

These entries would be queryable through the BPI Abstraction Layer APIs and would enable the system of record to quickly fetch the history of a single state object by providing only the unique state object ID.

When visualising the trustmesh for a single business transaction, it might look something like the following:

   
										     Sales Order V2 - Approved -------- Invoice V1 Approved
                                                    |
                                                    |
                                                    | 
	   Purchase Order v2 - Approved -------- Sales Order V1 - Declined
				|
			    |
				|
	   Purchase Order v1 - Declined									      		

In the graph, each node (or step) could then be enriched with metadata relevant to the business case (i.e relevant ids, rejection reasons etc.) This would enable consumers (i.e. systems of record) to build compelling and useful visualizations of the baselined business objects and  the processes around them.


# Rationale

We decided to propose extension of the transaction entity as it represents the place where the originator provides the already generated transaction id and has the option to attach additional system of record specific information (i.e. state object id). We propose to use the BPI storage to store the new entries as this same storage is being used to store state objects and their histories.

# Backwards Compatibility
None

# Test Cases
TBD

# Reference Implementation
Trustmesh implementation in Baseledger Lakewood https://github.com/Baseledger/baseledger-lakewood

# Security Considerations

There are no identified security implications to the proposed change. 
The data stored in the BPI is encrypted and is the same as the data that was already present there, with the addition of the unique state object id. 


---
Copyright
Copyright and related rights waived via CC0-Universal.

(This template adapted from the EIP template at https://github.com/ethereum/EIPs/