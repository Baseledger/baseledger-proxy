import { txClient, queryClient } from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'

import { SynchronizationFeedback } from "./module/types/trustmesh/SynchronizationFeedback"
import { SynchronizationRequest } from "./module/types/trustmesh/SynchronizationRequest"


async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
}

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
		structure.fields.push(field)
	}
	return structure
}

const getDefaultState = () => {
	return {
        SynchronizationFeedback: {},
        SynchronizationFeedbackAll: {},
        SynchronizationRequest: {},
        SynchronizationRequestAll: {},
        
        _Structure: {
            SynchronizationFeedback: getStructure(SynchronizationFeedback.fromPartial({})),
            SynchronizationRequest: getStructure(SynchronizationRequest.fromPartial({})),
            
		},
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(subscription)
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(subscription)
		}
	},
	getters: {
        getSynchronizationFeedback: (state) => (params = {}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SynchronizationFeedback[JSON.stringify(params)] ?? {}
		},
        getSynchronizationFeedbackAll: (state) => (params = {}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SynchronizationFeedbackAll[JSON.stringify(params)] ?? {}
		},
        getSynchronizationRequest: (state) => (params = {}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SynchronizationRequest[JSON.stringify(params)] ?? {}
		},
        getSynchronizationRequestAll: (state) => (params = {}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.SynchronizationRequestAll[JSON.stringify(params)] ?? {}
		},
        
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('init')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach((subscription) => {
				dispatch(subscription.action, subscription.payload)
			})
		},
		async QuerySynchronizationFeedback({ commit, rootGetters, getters }, { options: { subscribe = false , all = false}, params: {...key}, query=null }) {
			try {
				
				let value = query?(await (await initQueryClient(rootGetters)).querySynchronizationFeedback( key.id,  query)).data:(await (await initQueryClient(rootGetters)).querySynchronizationFeedback( key.id )).data
				
				commit('QUERY', { query: 'SynchronizationFeedback', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySynchronizationFeedback', payload: { options: { all }, params: {...key},query }})
				return getters['getSynchronizationFeedback']( { params: {...key}, query}) ?? {}
			} catch (e) {
				console.error(new SpVuexError('QueryClient:QuerySynchronizationFeedback', 'API Node Unavailable. Could not perform query.'))
				return {}
			}
		},
		async QuerySynchronizationFeedbackAll({ commit, rootGetters, getters }, { options: { subscribe = false , all = false}, params: {...key}, query=null }) {
			try {
				
				let value = query?(await (await initQueryClient(rootGetters)).querySynchronizationFeedbackAll( query)).data:(await (await initQueryClient(rootGetters)).querySynchronizationFeedbackAll()).data
				
				while (all && (<any> value).pagination && (<any> value).pagination.nextKey!=null) {
					let next_values=(await (await initQueryClient(rootGetters)).querySynchronizationFeedbackAll({...query, 'pagination.key':(<any> value).pagination.nextKey})).data
					for (let prop of Object.keys(next_values)) {
						if (Array.isArray(next_values[prop])) {
							value[prop]=[...value[prop], ...next_values[prop]]
						}else{
							value[prop]=next_values[prop]
						}
					}
				}
				
				commit('QUERY', { query: 'SynchronizationFeedbackAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySynchronizationFeedbackAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSynchronizationFeedbackAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				console.error(new SpVuexError('QueryClient:QuerySynchronizationFeedbackAll', 'API Node Unavailable. Could not perform query.'))
				return {}
			}
		},
		async QuerySynchronizationRequest({ commit, rootGetters, getters }, { options: { subscribe = false , all = false}, params: {...key}, query=null }) {
			try {
				
				let value = query?(await (await initQueryClient(rootGetters)).querySynchronizationRequest( key.id,  query)).data:(await (await initQueryClient(rootGetters)).querySynchronizationRequest( key.id )).data
				
				commit('QUERY', { query: 'SynchronizationRequest', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySynchronizationRequest', payload: { options: { all }, params: {...key},query }})
				return getters['getSynchronizationRequest']( { params: {...key}, query}) ?? {}
			} catch (e) {
				console.error(new SpVuexError('QueryClient:QuerySynchronizationRequest', 'API Node Unavailable. Could not perform query.'))
				return {}
			}
		},
		async QuerySynchronizationRequestAll({ commit, rootGetters, getters }, { options: { subscribe = false , all = false}, params: {...key}, query=null }) {
			try {
				
				let value = query?(await (await initQueryClient(rootGetters)).querySynchronizationRequestAll( query)).data:(await (await initQueryClient(rootGetters)).querySynchronizationRequestAll()).data
				
				while (all && (<any> value).pagination && (<any> value).pagination.nextKey!=null) {
					let next_values=(await (await initQueryClient(rootGetters)).querySynchronizationRequestAll({...query, 'pagination.key':(<any> value).pagination.nextKey})).data
					for (let prop of Object.keys(next_values)) {
						if (Array.isArray(next_values[prop])) {
							value[prop]=[...value[prop], ...next_values[prop]]
						}else{
							value[prop]=next_values[prop]
						}
					}
				}
				
				commit('QUERY', { query: 'SynchronizationRequestAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QuerySynchronizationRequestAll', payload: { options: { all }, params: {...key},query }})
				return getters['getSynchronizationRequestAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				console.error(new SpVuexError('QueryClient:QuerySynchronizationRequestAll', 'API Node Unavailable. Could not perform query.'))
				return {}
			}
		},
		
		async sendMsgCreateSynchronizationRequest({ rootGetters }, { value, fee, memo }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgCreateSynchronizationRequest(value)
				const result = await (await initTxClient(rootGetters)).signAndBroadcast([msg], {fee: { amount: fee, 
  gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgCreateSynchronizationRequest:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateSynchronizationRequest:Send', 'Could not broadcast Tx.')
				}
			}
		},
		async sendMsgUpdateSynchronizationRequest({ rootGetters }, { value, fee, memo }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgUpdateSynchronizationRequest(value)
				const result = await (await initTxClient(rootGetters)).signAndBroadcast([msg], {fee: { amount: fee, 
  gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgUpdateSynchronizationRequest:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateSynchronizationRequest:Send', 'Could not broadcast Tx.')
				}
			}
		},
		async sendMsgCreateSynchronizationFeedback({ rootGetters }, { value, fee, memo }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgCreateSynchronizationFeedback(value)
				const result = await (await initTxClient(rootGetters)).signAndBroadcast([msg], {fee: { amount: fee, 
  gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgCreateSynchronizationFeedback:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateSynchronizationFeedback:Send', 'Could not broadcast Tx.')
				}
			}
		},
		async sendMsgDeleteSynchronizationFeedback({ rootGetters }, { value, fee, memo }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgDeleteSynchronizationFeedback(value)
				const result = await (await initTxClient(rootGetters)).signAndBroadcast([msg], {fee: { amount: fee, 
  gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgDeleteSynchronizationFeedback:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeleteSynchronizationFeedback:Send', 'Could not broadcast Tx.')
				}
			}
		},
		async sendMsgDeleteSynchronizationRequest({ rootGetters }, { value, fee, memo }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgDeleteSynchronizationRequest(value)
				const result = await (await initTxClient(rootGetters)).signAndBroadcast([msg], {fee: { amount: fee, 
  gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgDeleteSynchronizationRequest:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeleteSynchronizationRequest:Send', 'Could not broadcast Tx.')
				}
			}
		},
		async sendMsgUpdateSynchronizationFeedback({ rootGetters }, { value, fee, memo }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgUpdateSynchronizationFeedback(value)
				const result = await (await initTxClient(rootGetters)).signAndBroadcast([msg], {fee: { amount: fee, 
  gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgUpdateSynchronizationFeedback:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateSynchronizationFeedback:Send', 'Could not broadcast Tx.')
				}
			}
		},
		
		async MsgCreateSynchronizationRequest({ rootGetters }, { value }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgCreateSynchronizationRequest(value)
				return msg
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgCreateSynchronizationRequest:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateSynchronizationRequest:Create', 'Could not create message.')
				}
			}
		},
		async MsgUpdateSynchronizationRequest({ rootGetters }, { value }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgUpdateSynchronizationRequest(value)
				return msg
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgUpdateSynchronizationRequest:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateSynchronizationRequest:Create', 'Could not create message.')
				}
			}
		},
		async MsgCreateSynchronizationFeedback({ rootGetters }, { value }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgCreateSynchronizationFeedback(value)
				return msg
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgCreateSynchronizationFeedback:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateSynchronizationFeedback:Create', 'Could not create message.')
				}
			}
		},
		async MsgDeleteSynchronizationFeedback({ rootGetters }, { value }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgDeleteSynchronizationFeedback(value)
				return msg
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgDeleteSynchronizationFeedback:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeleteSynchronizationFeedback:Create', 'Could not create message.')
				}
			}
		},
		async MsgDeleteSynchronizationRequest({ rootGetters }, { value }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgDeleteSynchronizationRequest(value)
				return msg
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgDeleteSynchronizationRequest:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeleteSynchronizationRequest:Create', 'Could not create message.')
				}
			}
		},
		async MsgUpdateSynchronizationFeedback({ rootGetters }, { value }) {
			try {
				const msg = await (await initTxClient(rootGetters)).msgUpdateSynchronizationFeedback(value)
				return msg
			} catch (e) {
				if (e.toString()=='wallet is required') {
					throw new SpVuexError('TxClient:MsgUpdateSynchronizationFeedback:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateSynchronizationFeedback:Create', 'Could not create message.')
				}
			}
		},
		
	}
}
