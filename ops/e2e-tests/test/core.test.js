var request = require('supertest');
const chai = require("chai");
const expect = chai.expect;
chai.use(require('chai-uuid'));

var alice_proxy_app_url = 'http://ub:ub321@localhost:8081';
var alice_blockchain_app_url = 'http://ub:ub321@localhost:1317';

var bob_proxy_app_url = 'http://ub:ub321@localhost:8082';
var bob_blockchain_app_url = 'http://ub:ub321@localhost:1318';

var test_workgroup_id = "734276bc-4adc-4621-acf8-ac66dc91cb27";
var alice_organization_id = "d45c9b93-3eef-4993-add6-aa1c84d17eea";
var bob_organization_id = "969e989c-bb61-4180-928c-0d48afd8c6a3";

const TEST_TIMEOUT = 30000;
const BLOCK_TIME_SLEEP_DELAY = 6000;

const sleep = (ms) => {
  return new Promise(resolve => setTimeout(resolve, ms));
};

describe('Setup orgs and workgroup', function () {
  it('Given Alice and Bob stacks, When both proxy apps triggered with organization and workgroup administration, it returns Ok', async function () {
    // get existing alice participation
    var getAliceParticipationResponse = await request(alice_proxy_app_url)
      .get(`/workgroup/${test_workgroup_id}/participation`)
      .send()
      .expect(200);
    
    var aliceParticipationId = JSON.parse(getAliceParticipationResponse.text)[0].id;


    // delete exisiting alice participation
    await request(alice_proxy_app_url)
      .delete(`/workgroup/${test_workgroup_id}/participation/${aliceParticipationId}`)
      .expect(204)

    var getAliceParticipationResponse = await request(bob_proxy_app_url)
      .get(`/workgroup/${test_workgroup_id}/participation`)
      .send()
      .expect(200);
    
    var aliceParticipationId = JSON.parse(getAliceParticipationResponse.text)[0].id;


    // delete exisiting alice participation
    await request(bob_proxy_app_url)
      .delete(`/workgroup/${test_workgroup_id}/participation/${aliceParticipationId}`)
      .expect(204)

    // get existing bob participation
    var getBobParticipationResponse = await request(bob_proxy_app_url)
      .get(`/workgroup/${test_workgroup_id}/participation`)
      .send()
      .expect(200);
    
    var bobParticipationId = JSON.parse(getBobParticipationResponse.text)[0].id;

    // delete exisiting bob participation
    await request(bob_proxy_app_url)
      .delete(`/workgroup/${test_workgroup_id}/participation/${bobParticipationId}`)
      .expect(204)
    
    // get existing bob participation
    var getBobParticipationResponse = await request(alice_proxy_app_url)
      .get(`/workgroup/${test_workgroup_id}/participation`)
      .send()
      .expect(200);
    
    var bobParticipationId = JSON.parse(getBobParticipationResponse.text)[0].id;

    // delete exisiting bob participation
    await request(alice_proxy_app_url)
      .delete(`/workgroup/${test_workgroup_id}/participation/${bobParticipationId}`)
      .expect(204)

    // create new alice participation
    const createAliceParticipationDto = {
      organization_id: alice_organization_id,
      organization_endpoint: "host.docker.internal:4222",
      organization_token: "testToken1"
    }

    await request(alice_proxy_app_url)
      .post(`/workgroup/${test_workgroup_id}/participation`)
      .send(createAliceParticipationDto)
      .expect(200);
    
    await request(bob_proxy_app_url)
      .post(`/workgroup/${test_workgroup_id}/participation`)
      .send(createAliceParticipationDto)
      .expect(200);


    // create new bob participation
    const createBobParticipationDto = {
      organization_id: bob_organization_id,
      organization_endpoint: "host.docker.internal:4223",
      organization_token: "testToken1"
    }
  
    await request(bob_proxy_app_url)
      .post(`/workgroup/${test_workgroup_id}/participation`)
      .send(createBobParticipationDto)
      .expect(200);
    
    await request(alice_proxy_app_url)
      .post(`/workgroup/${test_workgroup_id}/participation`)
      .send(createBobParticipationDto)
      .expect(200);
  });
});

describe('Send Suggestion and Feedback', function () {
  it('Given Alice and Bob stacks, When Alice proxy app is triggered with send suggestion and responds with feedback, it returns Ok with transaction hash and Alice and Bob nodes have the baseledger transactions', async function () {
    this.timeout(TEST_TIMEOUT + 20000);
    
    // SUGGESTION PART
    // Arrange
    const createSuggestionDto = {
      workgroup_id: test_workgroup_id,
      recipient: bob_organization_id,
      workstep_type: "FinalWorkstep",
      business_object_type: "PurchaseOrder",
      business_object_id: "169f104f-980e-42bb-a128-73daf259bc39",
      business_object_json: "{\"PurchaseOrderID\":\"PO123\",\"Currency\":\"EUR\",\"Amount\":\"200\"}"
    }

    // Act
    var createSuggestionResponse = await request(alice_proxy_app_url)
      .post('/suggestion')
      .send(createSuggestionDto)
      .expect(200);

    // Assert
    expect(createSuggestionResponse.body).not.to.be.undefined;
    console.log(createSuggestionResponse.body);
    
    var responseJSON = createSuggestionResponse.body
    console.log(responseJSON);

    const aliceSuggestionWorkflowId = responseJSON["workflow_id"];

    expect(aliceSuggestionWorkflowId).not.to.be.undefined;
    expect(aliceSuggestionWorkflowId).to.be.a.uuid();

    expect(responseJSON["workstep_id"]).not.to.be.undefined;
    expect(responseJSON["workstep_id"]).to.be.a.uuid();

    const aliceSuggestionBboid = responseJSON["baseledger_business_object_id"];

    expect(aliceSuggestionBboid).not.to.be.undefined;
    expect(aliceSuggestionBboid).to.be.a.uuid();

    expect(responseJSON["transaction_hash"]).not.to.be.undefined;
    expect(responseJSON["transaction_hash"]).to.have.length(64);

    expect(responseJSON["error"]).not.to.be.undefined;
    expect(responseJSON["error"]).to.be.empty;

    console.log(`WAITING ${BLOCK_TIME_SLEEP_DELAY}ms FOR A NEW BLOCK`);
    await sleep(BLOCK_TIME_SLEEP_DELAY);
    
    console.log(`WAITING SOME MORE FOR WORKER TO UPDATE TRUSTMESH STATUS`);
    await sleep(BLOCK_TIME_SLEEP_DELAY);

    var queryAliceTransactionsResponse = await request(alice_blockchain_app_url)
      .get('/unibrightio/baseledger/baseledger/BaseledgerTransaction')
      .expect(200);

    var payload = JSON.parse(queryAliceTransactionsResponse.text);
    console.log("queryAliceTransactionsResponse " + JSON.stringify(payload));
    expect(payload.BaseledgerTransaction).not.to.be.undefined;
    expect(payload.BaseledgerTransaction).to.have.length.above(0);

    var queryBobTransactionsResponse = await request(bob_blockchain_app_url)
      .get('/unibrightio/baseledger/baseledger/BaseledgerTransaction')
      .expect(200);

    var payload = JSON.parse(queryBobTransactionsResponse.text);
    console.log("queryBobTransactionsResponse " + JSON.stringify(payload));

    expect(payload.BaseledgerTransaction).not.to.be.undefined;
    expect(payload.BaseledgerTransaction).to.have.length.above(0);

    var getAliceTrustmeshResponse = await request(alice_proxy_app_url)
      .get('/trustmeshes')
      .expect(200);
    
    var trustmesheshPayload = JSON.parse(getAliceTrustmeshResponse.text);
    console.log("getAliceTrustmeshResponse " + JSON.stringify(trustmesheshPayload));

    expect(trustmesheshPayload[0]["Entries"][0].EntryType).to.equal("SuggestionSent");
    expect(trustmesheshPayload[0]["Entries"][0].CommitmentState).to.equal("COMMITTED");

    var getBobTrustmeshResponse = await request(bob_proxy_app_url)
      .get('/trustmeshes')
      .expect(200);
  
    var trustmesheshPayload = JSON.parse(getBobTrustmeshResponse.text);
    console.log("getBobTrustmeshResponse " + JSON.stringify(trustmesheshPayload));
    expect(trustmesheshPayload[0]["Entries"][0].EntryType).to.equal("SuggestionReceived");
    expect(trustmesheshPayload[0]["Entries"][0].CommitmentState).to.equal("COMMITTED");

    // FEEDBACK PART
    // Arrange
    const createFeedbackDto = {
      baseledger_business_object_id: aliceSuggestionBboid,
      approved: true,
      feedback_message: ""
    }

    // Act
    var createFeedbackResponse = await request(bob_proxy_app_url)
      .post('/feedback')
      .send(createFeedbackDto)
      .expect(200);

    // Assert
    expect(createFeedbackResponse.body).not.to.be.undefined;

    var responseJSON = createFeedbackResponse.body
    console.log(responseJSON);

    const bobFeedbackWorkflowId = responseJSON["workflow_id"];

    expect(bobFeedbackWorkflowId).not.to.be.undefined;
    expect(bobFeedbackWorkflowId).to.be.a.uuid();

    expect(responseJSON["workstep_id"]).not.to.be.undefined;
    expect(responseJSON["workstep_id"]).to.be.a.uuid();

    const bobFeedbackBboid = responseJSON["baseledger_business_object_id"];

    expect(bobFeedbackBboid).not.to.be.undefined;
    expect(bobFeedbackBboid).to.be.equal(aliceSuggestionBboid);

    expect(responseJSON["transaction_hash"]).not.to.be.undefined;
    expect(responseJSON["transaction_hash"]).to.have.length(64);

    expect(responseJSON["error"]).not.to.be.undefined;
    expect(responseJSON["error"]).to.be.empty;

    console.log(`WAITING ${BLOCK_TIME_SLEEP_DELAY}ms FOR A NEW BLOCK`);
    await sleep(BLOCK_TIME_SLEEP_DELAY);
    
    console.log(`WAITING SOME MORE FOR WORKER TO UPDATE TRUSTMESH STATUS`);
    await sleep(BLOCK_TIME_SLEEP_DELAY);

    var queryAliceTransactionsResponse = await request(alice_blockchain_app_url)
      .get('/unibrightio/baseledger/baseledger/BaseledgerTransaction')
      .expect(200);

    var payload = JSON.parse(queryAliceTransactionsResponse.text);
    expect(payload.BaseledgerTransaction).not.to.be.undefined;
    expect(payload.BaseledgerTransaction).to.have.length.above(0);

    var queryBobTransactionsResponse = await request(bob_blockchain_app_url)
      .get('/unibrightio/baseledger/baseledger/BaseledgerTransaction')
      .expect(200);

    var payload = JSON.parse(queryBobTransactionsResponse.text);
    expect(payload.BaseledgerTransaction).not.to.be.undefined;
    expect(payload.BaseledgerTransaction).to.have.length.above(0);

    var getAliceTrustmeshResponse = await request(alice_proxy_app_url)
      .get('/trustmeshes')
      .expect(200);

    var trustmesheshPayload = JSON.parse(getAliceTrustmeshResponse.text);
    console.log("getAliceTrustmeshResponse 2 " + JSON.stringify(trustmesheshPayload));

    expect(trustmesheshPayload[0]["Entries"][1].EntryType).to.equal("FeedbackReceived");
    expect(trustmesheshPayload[0]["Entries"][1].CommitmentState).to.equal("COMMITTED");

    var getBobTrustmeshResponse = await request(bob_proxy_app_url)
      .get('/trustmeshes')
      .expect(200);

    var trustmesheshPayload = JSON.parse(getBobTrustmeshResponse.text);
    console.log("getBobTrustmeshResponse 2 " + JSON.stringify(trustmesheshPayload));
    expect(trustmesheshPayload[0]["Entries"][1].EntryType).to.equal("FeedbackSent");
    expect(trustmesheshPayload[0]["Entries"][1].CommitmentState).to.equal("COMMITTED");
  });
});