var request = require('supertest');
var uuid = require("uuid");
var expect = require('chai').expect;

var alice_proxy_app_url = 'http://localhost:8081';
var alice_blockchain_app_url = 'http://localhost:1317';

var bob_proxy_url = 'http://localhost:8082';
var bob_blockchain_app_url = 'http://localhost:1318';

var test_workgroup_id = "734276bc-4adc-4621-acf8-ac66dc91cb27";
var bob_organization_id = "969e989c-bb61-4180-928c-0d48afd8c6a3";

const BLOCK_TIME_SLEEP_DELAY = 6000;

const sleep = (ms) => {
  return new Promise(resolve => setTimeout(resolve, ms));
};

describe('Send Suggestion', function () {
  it('Given Alice and Bob stacks, When Alice proxy app is triggered with send suggestion, it returns Ok with transaction hash and Alice and Bob nodes have the baseledger transaction', async function () {
    this.timeout(BLOCK_TIME_SLEEP_DELAY + 2000);
    
    // Arrange
    const createSuggestionDto = {
      workgroup_id: test_workgroup_id,
      recipient: bob_organization_id,
      workstep_type: "FinalWorkstep",
      business_object_type: "PurchaseOrder",
      baseledger_business_object_id: "169f104f-980e-42bb-a128-73daf259bc39",
      business_object_json: "{\"PurchaseOrderID\":\"PO123\",\"Currency\":\"EUR\",\"MaterialID\":\"4711\",\"Quantity\":3,\"SinglePrice\":2.5,\"TotalPrice\":7.5,\"OrderItems\":[{\"ItemID\":20,\"ItemMaterialID\":\"FRT45098\",\"ItemQuantity\":70,\"Texts\":[\"text569hngf3ei\",\"j908j9j9j\",\"2340rfjmn2roicvn\"]},{\"ItemID\":20,\"ItemMaterialID\":\"FRT45098\",\"ItemQuantity\":70,\"Texts\":[\"text569hngf3ei\",\"j908j9j9j\",\"2340rfjmn2roicvn\"]},{\"ItemID\":20,\"ItemMaterialID\":\"FRT45098\",\"ItemQuantity\":70,\"Texts\":[\"text569hngf3ei\",\"j908j9j9j\",\"2340rfjmn2roicvn\"]},{\"ItemID\":20,\"ItemMaterialID\":\"FRT45098\",\"ItemQuantity\":70,\"Texts\":[\"text569hngf3ei\",\"j908j9j9j\",\"2340rfjmn2roicvn\"]},{\"ItemID\":20,\"ItemMaterialID\":\"FRT45098\",\"ItemQuantity\":70,\"Texts\":[\"text569hngf3ei\",\"j908j9j9j\",\"2340rfjmn2roicvn\"]}]}",
      referenced_baseledger_business_object_id: "",
      referenced_baseledger_transaction_id: ""
    }

    // Act
    var createSuggestionResponse = await request(alice_proxy_app_url)
      .post('/suggestion')
      .send(createSuggestionDto)
      .expect(200);

    // Assert
    expect(createSuggestionResponse.body).not.to.be.undefined;
    expect(createSuggestionResponse.body).to.have.length(64);

    console.log(`WAITING ${BLOCK_TIME_SLEEP_DELAY}ms FOR A NEW BLOCK`);
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
  });
});