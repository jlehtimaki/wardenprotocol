// Generated by Ignite ignite.com/cli
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { Api } from "./rest";
import { MsgCreatePermanentLockedAccountResponse } from "./types/cosmos/vesting/v1beta1/tx";
import { BaseVestingAccount } from "./types/cosmos/vesting/v1beta1/vesting";
import { ContinuousVestingAccount } from "./types/cosmos/vesting/v1beta1/vesting";
import { Period } from "./types/cosmos/vesting/v1beta1/vesting";
import { MsgCreatePermanentLockedAccount } from "./types/cosmos/vesting/v1beta1/tx";
import { PermanentLockedAccount } from "./types/cosmos/vesting/v1beta1/vesting";
import { MsgCreateVestingAccount } from "./types/cosmos/vesting/v1beta1/tx";
import { MsgCreateVestingAccountResponse } from "./types/cosmos/vesting/v1beta1/tx";
import { MsgCreatePeriodicVestingAccount } from "./types/cosmos/vesting/v1beta1/tx";
import { PeriodicVestingAccount } from "./types/cosmos/vesting/v1beta1/vesting";
import { MsgCreatePeriodicVestingAccountResponse } from "./types/cosmos/vesting/v1beta1/tx";
import { DelayedVestingAccount } from "./types/cosmos/vesting/v1beta1/vesting";
export { MsgCreatePermanentLockedAccountResponse, BaseVestingAccount, ContinuousVestingAccount, Period, MsgCreatePermanentLockedAccount, PermanentLockedAccount, MsgCreateVestingAccount, MsgCreateVestingAccountResponse, MsgCreatePeriodicVestingAccount, PeriodicVestingAccount, MsgCreatePeriodicVestingAccountResponse, DelayedVestingAccount };
export const registry = new Registry(msgTypes);
function getStructure(template) {
    const structure = { fields: [] };
    for (let [key, value] of Object.entries(template)) {
        let field = { name: key, type: typeof value };
        structure.fields.push(field);
    }
    return structure;
}
const defaultFee = {
    amount: [],
    gas: "200000",
};
export const txClient = ({ signer, prefix, addr } = { addr: "http://localhost:26657", prefix: "cosmos" }) => {
    return {
        async sendMsgCreatePermanentLockedAccountResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCreatePermanentLockedAccountResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.msgCreatePermanentLockedAccountResponse({ value: MsgCreatePermanentLockedAccountResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCreatePermanentLockedAccountResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendBaseVestingAccount({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendBaseVestingAccount: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.baseVestingAccount({ value: BaseVestingAccount.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendBaseVestingAccount: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendContinuousVestingAccount({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendContinuousVestingAccount: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.continuousVestingAccount({ value: ContinuousVestingAccount.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendContinuousVestingAccount: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendPeriod({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendPeriod: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.period({ value: Period.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendPeriod: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgCreatePermanentLockedAccount({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCreatePermanentLockedAccount: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.msgCreatePermanentLockedAccount({ value: MsgCreatePermanentLockedAccount.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCreatePermanentLockedAccount: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendPermanentLockedAccount({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendPermanentLockedAccount: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.permanentLockedAccount({ value: PermanentLockedAccount.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendPermanentLockedAccount: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgCreateVestingAccount({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCreateVestingAccount: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.msgCreateVestingAccount({ value: MsgCreateVestingAccount.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCreateVestingAccount: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgCreateVestingAccountResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCreateVestingAccountResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.msgCreateVestingAccountResponse({ value: MsgCreateVestingAccountResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCreateVestingAccountResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgCreatePeriodicVestingAccount({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCreatePeriodicVestingAccount: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.msgCreatePeriodicVestingAccount({ value: MsgCreatePeriodicVestingAccount.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCreatePeriodicVestingAccount: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendPeriodicVestingAccount({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendPeriodicVestingAccount: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.periodicVestingAccount({ value: PeriodicVestingAccount.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendPeriodicVestingAccount: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendMsgCreatePeriodicVestingAccountResponse({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendMsgCreatePeriodicVestingAccountResponse: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.msgCreatePeriodicVestingAccountResponse({ value: MsgCreatePeriodicVestingAccountResponse.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendMsgCreatePeriodicVestingAccountResponse: Could not broadcast Tx: ' + e.message);
            }
        },
        async sendDelayedVestingAccount({ value, fee, memo }) {
            if (!signer) {
                throw new Error('TxClient:sendDelayedVestingAccount: Unable to sign Tx. Signer is not present.');
            }
            try {
                const { address } = (await signer.getAccounts())[0];
                const signingClient = await SigningStargateClient.connectWithSigner(addr, signer, { registry });
                let msg = this.delayedVestingAccount({ value: DelayedVestingAccount.fromPartial(value) });
                return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo);
            }
            catch (e) {
                throw new Error('TxClient:sendDelayedVestingAccount: Could not broadcast Tx: ' + e.message);
            }
        },
        msgCreatePermanentLockedAccountResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccountResponse", value: MsgCreatePermanentLockedAccountResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCreatePermanentLockedAccountResponse: Could not create message: ' + e.message);
            }
        },
        baseVestingAccount({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.BaseVestingAccount", value: BaseVestingAccount.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:BaseVestingAccount: Could not create message: ' + e.message);
            }
        },
        continuousVestingAccount({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.ContinuousVestingAccount", value: ContinuousVestingAccount.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:ContinuousVestingAccount: Could not create message: ' + e.message);
            }
        },
        period({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.Period", value: Period.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:Period: Could not create message: ' + e.message);
            }
        },
        msgCreatePermanentLockedAccount({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccount", value: MsgCreatePermanentLockedAccount.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCreatePermanentLockedAccount: Could not create message: ' + e.message);
            }
        },
        permanentLockedAccount({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.PermanentLockedAccount", value: PermanentLockedAccount.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:PermanentLockedAccount: Could not create message: ' + e.message);
            }
        },
        msgCreateVestingAccount({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.MsgCreateVestingAccount", value: MsgCreateVestingAccount.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCreateVestingAccount: Could not create message: ' + e.message);
            }
        },
        msgCreateVestingAccountResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse", value: MsgCreateVestingAccountResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCreateVestingAccountResponse: Could not create message: ' + e.message);
            }
        },
        msgCreatePeriodicVestingAccount({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount", value: MsgCreatePeriodicVestingAccount.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCreatePeriodicVestingAccount: Could not create message: ' + e.message);
            }
        },
        periodicVestingAccount({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.PeriodicVestingAccount", value: PeriodicVestingAccount.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:PeriodicVestingAccount: Could not create message: ' + e.message);
            }
        },
        msgCreatePeriodicVestingAccountResponse({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse", value: MsgCreatePeriodicVestingAccountResponse.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:MsgCreatePeriodicVestingAccountResponse: Could not create message: ' + e.message);
            }
        },
        delayedVestingAccount({ value }) {
            try {
                return { typeUrl: "/cosmos.vesting.v1beta1.DelayedVestingAccount", value: DelayedVestingAccount.fromPartial(value) };
            }
            catch (e) {
                throw new Error('TxClient:DelayedVestingAccount: Could not create message: ' + e.message);
            }
        },
    };
};
export const queryClient = ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseURL: addr });
};
class SDKModule {
    constructor(client) {
        this.registry = [];
        this.query = queryClient({ addr: client.env.apiURL });
        this.updateTX(client);
        this.structure = {};
        client.on('signer-changed', (signer) => {
            this.updateTX(client);
        });
    }
    updateTX(client) {
        const methods = txClient({
            signer: client.signer,
            addr: client.env.rpcURL,
            prefix: client.env.prefix ?? "cosmos",
        });
        this.tx = methods;
        for (let m in methods) {
            this.tx[m] = methods[m].bind(this.tx);
        }
    }
}
;
const IgntModule = (test) => {
    return {
        module: {
            CosmosVestingV1Beta1: new SDKModule(test)
        },
        registry: msgTypes
    };
};
export default IgntModule;
