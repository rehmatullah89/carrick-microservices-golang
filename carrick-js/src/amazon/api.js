import {ClIdTag} from "./clid";
import {Cache} from "./cache";

const Api = {
    baseUrl: process.env.MIX_API_URL,

    getTag(publisherHash) {
        let url = new URL(`${Api.baseUrl}/get-tag/${publisherHash}`);
        let params = {
            t_type: ClIdTag._tagType ? ClIdTag._tagType.toLowerCase() : ''
        };

        if (ClIdTag._tagType === 'Organic') {
            Object.assign(params,{
                r: Cache.get('r'),
                a_u: window.location.href
            });
        }

        url.search = new URLSearchParams(params);

        return fetch(url)
            .then(response => response.json())
            .then(result => {
                if (!result.status || !result.data) {
                    throw "Request is not successful.";
                }

                return result.data
            });
    },
    sendVisit(publisherHash, {t, ci_t, ci_v, r, a_u, u_agent}) {
        const data = {t, ci_t, ci_v, r, a_u, u_agent};

        return Api.sendApi(`${Api.baseUrl}/visit/${publisherHash}`, data);
    },
    sendTracking(publisherHash, {t, ci_t, ci_v, r, a_u, c_u, u_agent}) {
        const data = {t, ci_t, ci_v, r, a_u, c_u, u_agent};

        return Api.sendApi(`${Api.baseUrl}/send-tracking/${publisherHash}`, data);
    },
    // FIXME: In the possible future.
    //  This method can return an unexpected value depending on the browser.
    sendApi(url, data) {
        if (navigator.sendBeacon) {
            return navigator.sendBeacon(url, JSON.stringify(data));
        }

        const req = new XMLHttpRequest();

        req.open('POST', url, false);
        req.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        req.send(JSON.stringify(data));

        return req.status === 200;
    }
};

export {Api}