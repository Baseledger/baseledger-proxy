// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import UnibrightioBaseledgerUnibrightioBaseledgerBaseledger from './unibrightio/baseledger/unibrightio.baseledger.baseledger';
export default {
    UnibrightioBaseledgerUnibrightioBaseledgerBaseledger: load(UnibrightioBaseledgerUnibrightioBaseledgerBaseledger, 'unibrightio.baseledger.baseledger'),
};
function load(mod, fullns) {
    return function init(store) {
        const fullnsLevels = fullns.split('/');
        for (let i = 1; i < fullnsLevels.length; i++) {
            let ns = fullnsLevels.slice(0, i);
            if (!store.hasModule(ns)) {
                store.registerModule(ns, { namespaced: true });
            }
        }
        if (store.hasModule(fullnsLevels)) {
            throw new Error('Duplicate module name detected: ' + fullnsLevels.pop());
        }
        else {
            store.registerModule(fullnsLevels, mod);
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns + '/init', null, {
                        root: true
                    });
                }
            });
        }
    };
}
