import {Config} from "./config";
import {Cache} from "./cache";

const Helpers = {
    getReferer() {
        return Cache.get('r');
    },
    checkUrlParameterExists(param, url) {
        url = !url ? window.location.search : url;

        if (url.indexOf('?' + param + '=') != -1 || url.indexOf('&' + param + '=') != -1)
            return true;

        return false
    },

    replaceTagParam(urlString, newParam) {
        if (urlString.indexOf('tag=') !== -1) {
            return urlString.replace(/([?&]tag=)[^&]+/gi, '$1' + newParam);
        }

        let url = new URL(urlString);
        let search = url.search.substring(1);

        if (search) {
            search = '&' + search;
        }

        return url.origin + url.pathname + '?tag=' + newParam + '' + search;
    },
    getUrlParameter(sParam, url) {
        url = !url ? window.location.search : url;

        var sPageURL = url.substring(1),
            sURLVariables = sPageURL.split('&'),
            sParameterName,
            i;

        for (i = 0; i < sURLVariables.length; i++) {
            sParameterName = sURLVariables[i].split('=');

            if (sParameterName[0] === sParam) {
                return typeof sParameterName[1] === undefined ? true : decodeURIComponent(sParameterName[1]);
            }
        }
        return false;
    },
    _generateUUID() {
        const e = Date.now(),
            t = String.fromCharCode(Math.floor(10 * Math.random() + 97)),
            i = String.fromCharCode(Math.floor(10 * Math.random() + 97)),
            r = String.fromCharCode(Math.floor(10 * Math.random() + 97));
        return `fr${e}${t}${i}${r}`;
    },
    replaceSubIdParam(url, subId1) {

        if (url.href.indexOf('subId1=') !== -1) {
            return url.href.replace(/([?&]subId1=)[^&]+/gi, '$1' + subId1);
        }
        // let url = new URL(urlString);
        let search = url.search.substring(1);
        if (search) {
            search = '&' + search;
        }
        return url.origin + url.pathname + '?subId1=' + subId1 + '' + search;
    },
    replaceXcustParam(url, xcust) {

        if (url.href.indexOf('xcust=') !== -1) {
            return url.href.replace(/([?&]xcust=)[^&]+/gi, '$1' + xcust);
        }
        let search = url.search.substring(1);
        if (search) {
            search = '&' + search;
        }
        return url.origin + url.pathname + '?xcust=' + xcust + '' + search;
    },
    replaceU1Param(url, u1) {

        if (url.href.indexOf('u1=') !== -1) {
            return url.href.replace(/([?&]u1=)[^&]+/gi, '$1' + u1);
        }
        let search = url.search.substring(1);
        if (search) {
            search = '&' + search;
        }
        return url.origin + url.pathname + '?u1=' + u1 + '' + search;
    },
    transformUrls(newTag, anchors = []) {
        let url;
        let urls = !anchors.length ? document.querySelectorAll(Config.trackedSelector) : anchors;

        for (let i = 0; i < urls.length; i++) {
            url = urls[i];

            url.href = Helpers.replaceTagParam(url.href, newTag);
            url.classList.add('carrick-track');
        }
    },
    transformSkimlinksUrls() {
        let urls = document.querySelectorAll(Config.trackedSkimlinksSelectors)
        // var urls = document.querySelectorAll('a[href*="go.skimresources.com"]')
        for (let i = 0; i < urls.length; i++) {
            let url = urls[i];
            let custom_id = 'skmrvid:' + Helpers._generateUUID();
            url.href = Helpers.replaceXcustParam(url, custom_id);

            // let custom_id = Helpers.getUrlParameter('xcust', url.href)
            url.setAttribute('data-skim-tracking', custom_id)
            url.classList.add('skimlinks-track')
        }
    },
    transformImpactUrls() {
        // console.log(Config.trackedImpactSelectors)
        let urls = document.querySelectorAll(Config.trackedImpactSelectors)
        for (let i = 0; i < urls.length; i++) {
            let url = urls[i];
            let subId1 = 'imprvid:' + Helpers._generateUUID();

            url.href = Helpers.replaceSubIdParam(url, subId1);
            url.classList.add('impact-track')
            url.setAttribute('data-imp-tracking', subId1);
        }
    },
    transformNarrativUrls() {
        // console.log(Config.trackedNarrativSelectors)
        let urls = document.querySelectorAll(Config.trackedNarrativSelectors)
        for (let i = 0; i < urls.length; i++) {
            let url = urls[i];
            let U1 = 'nrvrvid:' + Helpers._generateUUID();

            url.href = Helpers.replaceU1Param(url,  U1);
            url.classList.add('narrativ-track')
            url.setAttribute('data-nrv-tracking', U1);
        }
    },
    delegateEvent(event, callback, targetSelector, parentElement) {
        parentElement.addEventListener(event, (e) => {
            const elem = e.target.closest(targetSelector);

            if (elem && callback instanceof Function) {
                callback.call(elem, e, elem);
            }
        });
    },
};

export {Helpers}