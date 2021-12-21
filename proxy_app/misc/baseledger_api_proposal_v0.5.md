# Baseledger Proxy API proposal



## Suggestion

**POST /suggestion**

body:
```json
{
    "workgroup_id": "", // Optional. GUID, represents the workgroup where the recipient is located. If ommitted, assumption is that recipient stated bellow has only one workgroup. 

    "recipient": "", // Mandatory. GUID, represents the recipient. We curently support only one recipient. 
    
    "workstep_type": "", // Optional. INITIAL, NEWVERSION, NEXTWORKSTEP, FINAL. Look bellow for rules related to this field and relation with baseledger_business_object_id.

    "workflow_id": "", // Mandatory if baseledger_business_object_id not provided and workstep NEWVERSION, NEXTWORKSTEP or FINAL. If baseledger_business_object_id is provided, it preceeds over this one. Maps to trustmesh_id internally.
    
    "baseledger_business_object_id": "", // Optional. Look bellow for rules related to this field and relation with workstep_type.

    "business_object_type": "", // Optional. Text field representing the type of the object in the SOR. Can be anything and we just store it.

    "business_object_id": "", // Mandatory. Text field representing unique id of the object in the SOR. Can be anything and we just store it.

    "business_object_json": "", // Mandatory. Json payload of the business object.

    "knowledge_limiters" // Optional. Used to hide away payload properties. Not relevant for now.
}
```

response(200 ok, 400 in case of processing error in the proxy):
```json
    {
        "workflow_id": "", // ID of the trustmesh where the bboid belongs
        "workstep_id": "", // Latest entry of the bboid in the trustmesh (i.e. suggestion for new version sent, waiting feedback)
        "baseledger_business_object_id": "", // newly generated in case INITIAL, NEXTWORKSTEP or FINAL suggestion. Same if NEWVERSION
        "transaction_hash": "", // Latest transaction hash of the relevant trustmesh entry
        "error": "" // string message in case of failure
    }
```

**Notes on workflow_id <-> bboid <-> workstep_type relation**

* If baseledger_business_object_id not provided and workflow_id not provided
    Default to INITIAL workstep_type
    Create a bboid, create suggestion sent trustmesh entry and return to the SOR.

* if bboid is provided and NEWVERSION workstep_type,
    find the latest trustmesh entry based on the bboid, verify trustmesh entry in correct state (i.e. not finalized, not waiting for feedback etc.), create a new suggestion with the same bboid and return bboid to SOR

* if bboid is provided and NEXTWORKSTEP workstep_type,
    this bboid references the previous workstep bussiness object.
    find the latest trustmesh entry based on the bboid, verify correct state (i.e. not finalized, not waiting for feedback etc.),
    create a new suggestion with a newly generated bboid as the bboid and the SOR provided as the referencedbboid and return to SOR

* if bboid is provided and FINAL workstep_type,
    this bboid references the previous workstep bussiness object.
    find the latest trustmesh entry based on the bboid, verify correct state (i.e. not finalized, not waiting for feedback etc.),
    create a new suggestion with a newly generated bboid as the bboid and the SOR provided as the referencedbboid and return to SOR

* If bboid not provided, workflow_id provided and workstep_type NEWVERSION, NEXTWORKSTEP, FINAL
    Find the latest trustmesh entry based on the workflow_id. If status ok, create a new suggestion 
    and return the bboid.


With this approach, the SOR always keeps track of a single bboid per object, and uses this as a reference to continue the workflow.
if there is a new object that needs to be created in the SOR as a continuation of the workflow (invoice from baselined purchase order) this object would
need to pick up the bboid from the previous object in order to be able to continue the same workflow with a new workstep.

Additionaly, SOR can just keep track of the workflow id and let proxy deal with the latest suggestion\feebdack

## Feedback

**POST /feedback**

body:

```json
{
     "workflow_id": "", // Mandatory if baseledger_business_object_id not provided. If baseledger_business_object_id is provided, it preceeds over this one. Maps to trustmesh_id internally.
     "baseledger_business_object_id": "", // // Mandatory. The referenced id must in be in at least one trustmesh entry and the entry has to be in the correct state to continue.
     "approved": true, // Mandatory. true, false.
     "feedback_message": "" // Optional. Text following the feedback 
}
```

response(200 ok, 400 in case of processing error in the proxy):

```json    
{
    "workflow_id": "", // ID of the trustmesh where the bboid belongs
    "workstep_id": "", // Latest entry of the bboid in the trustmesh (FeedbackSent trustmesh entry)
    "baseledger_business_object_id": "", // same bboid as the one sent in the request
    "transaction_hash": "", // Latest transaction hash of the relevant trustmesh entry
    "error": // string message in case of failure
}
```

## Webhook

**POST /webhook**

body:
```json  
{
    "url": "", // Mandatory. Url representing the sor endpoint to trigger. Params in the url must be defined with the following syntax {{param_name}}
    "url_params": [], // Optional. List of key values represting url parameter name (param_name) -> url parameter value field. Parameter value can be anything from the trustmesh entry properties listed bellow. Special params do not have to be listed here.
    "http_method": "", // Mandatory. PUT or POST
    "webhook_type": "", // Mandatory. 0 - Create object, 1 - Update status. 0 is triggered whenever there is a new suggestion. 1 is triggered for proxy replies or feedbacks.
    "auth_type": "", // Mandatory. 0 - None, 1 - Basic auth
    "auth_username": "", // Mandatory if auth_type > 0
    "auth_password": "", // Mandatory if auth_type > 0
    "xcsrf_url": "", // Optional. If provided, used to fetch the token and place it in header of every request. Provided auth_type will also be applied.
    "body_content_type": "", // Optional. JSON, XML or other format. JSON supported for now. Defaults to JSON if not provided.
    "body": "", // Mandatory. Body of the request. Params in the body must be defined with the following syntax {{param_name}}. Special params listed bellow.
    "body_params": "" // Optional. List of key values represting body parameter name -> body parameter value field Parameter value can be anything from the trustmesh entry properties listed bellow. Special params do not have to be listed here.
}
```
response(200 ok, 400 in case of processing error in the proxy):

```json    
{
    "webhook_id": "", // GUID of the newly created webhook if successful
    "error": "",  // string error if code != 200
}
```

**Note on trustmesh entry properties that can be used in url or body params**

Id
TendermintBlockId
TendermintTransactionId
TendermintTransactionTimestamp
EntryType
SenderOrgId
ReceiverOrgId
SenderOrg
ReceiverOrg
WorkgroupId
Workgroup
WorkstepType
BaseledgerTransactionType
BaseledgerTransactionId
ReferencedBaseledgerTransactionId
BusinessObjectType
BaseledgerBusinessObjectId
ReferencedBaseledgerBusinessObjectId
OffchainProcessMessageId
OffchainProcessMessage
CommitmentState
TransactionHash
TrustmeshId
SorBusinessObjectId

Usage example: 

    "url_params": [{"ParamName":"my_custom_param_name", "ParamValueField": "BaseledgerBusinessObjectId"}]

**Note on special params**

{{origin}} - empty string if comming from proxy, otherwise sender organization id
{{message}} - error message from proxy or feedback message from sender organization id
{{business_object_json_payload}} - if webhook_type = 0 (Incoming new suggestion for object creation), this holds the business object payload
{{approved}} - if webhook_type = 1 (Incoming feebdack for an existing object), this hold the feedback status true - Approved or false - Decline
{{organization_id}} - id of the organization stored in the proxy instance of the webhook creator

**GET /webhook**
lists all

**GET /webhook/{id}**
deletes given if exist

## Pooling

**GET /workflow/new**

return only trustmesh entries where suggestion received is the latest state

response(200 ok, 400 in case of processing error in the proxy):
```json
    [{
        "workflow_id": "", // ID of the new trustmesh
        "workstep_id": "", // Id of the latest trustmesh entry id
        "workstep_type": "", // INITIAL
        "baseledger_business_object_id": "", // bboid of the business object
        "business_object_json_payload": "", // Payload of the business object
    }]
```

**GET /workflow/lateststatebybboid/{baseledger_business_object_id}**

return latest trustmesh entry for a specific baseledger_business_object_id

response(200 ok, 400 in case of processing error in the proxy):
```json
    [{
        "workflow_id": "", // ID of the trustmesh
        "workstep_id": "", // Latest trustmesh entry id (i.e. suggestion sent, feedbackreceived)
        "workstep_type": "", // SUGGESTION, FEEDBACK, NEWVERSION, NEXT_WORKSTEP, FINAL
        "baseledger_business_object_id": "", // newly generated if SUGGESTION, NEXTWORKSTEP or FINAL suggestion. Same if NEWVERSION, FEEDBACK
        "business_object_json_payload": "", // Payload of the business object if workstep_type SUGGESTION, NEWVERSION, NEXTWORKSTEP
        "approved": true, // If workstep_type FEEDBACK, true for approved, false for declined
    }]
```