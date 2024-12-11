import {Helpers} from "./helpers";
import {Cache} from "./cache";
import {Api} from "./api";
import {Config} from "./config";

const ClIdTag = {
    _tagType: '',
    _publisherHash: null,

    get(publisherHash) {
        this._publisherHash = publisherHash;

        return new Promise((resolve, reject) => {
            ClIdTag._rememberReferrer();

            // url exists and cookie exists
            if (ClIdTag._resolvePaid(resolve)) {
                // TODO: Store known cache keys in config
                // TODO Remove this code
                Cache.remove('tag', document.location.pathname);
                Cache.remove('clid_t', document.location.pathname);
                Cache.remove('clid_v', document.location.pathname);

                return;
            }
            ClIdTag._resolveOrganic(resolve);
        });
    },
    handleVisit(publisherHash) {
        let clIdFromUrl = this._getClIdFromUrl();

        if (!clIdFromUrl) {
            clIdFromUrl = {
                type: Cache.get('clid_t'),
                value: Cache.get('clid_v')
            };
        }

        this._publisherHash = publisherHash;

        Api.sendVisit(publisherHash, {
            t: Cache.get('tag'),
            ci_t: clIdFromUrl.type,
            ci_v: clIdFromUrl.value,
            r: Helpers.getReferer(),
            a_u: window.location.href,
            u_agent: navigator.userAgent
        })
    },
    handleTracking(publisherHash) {
        const callback = (e, url) => {
            Api.sendTracking(
                publisherHash,
                {
                    t: Cache.get('tag'),
                    ci_t: Cache.get('clid_t'),
                    ci_v: Cache.get('clid_v'),
                    r: Helpers.getReferer(),
                    a_u: window.location.href,
                    c_u: url.href,
                    u_agent: navigator.userAgent
                }
            );
        };

        Helpers.delegateEvent('click', callback, '.carrick-track', document.body);
        ClIdTag._registerMiddleClick(callback);
    },
    handleSkimlinksTracking(publisherHash) {

        let clId = this._getClIdFromNonAmazon();
        const callback = (e, url) => {
            Api.sendTracking(
                publisherHash,
                {
                    t: url.getAttribute('data-skim-tracking'),
                    ci_t: clId.type,
                    ci_v: clId.value,
                    r: Helpers.getReferer(),
                    a_u: window.location.href,
                    c_u: url.href,
                    u_agent: navigator.userAgent
                }
            );
        };

        Helpers.delegateEvent('click', callback, '.skimlinks-track', document.body);
        ClIdTag._registerSkimlinksMiddleClick(callback);
    },
    handleImpactTracking(publisherHash) {
        let clId = this._getClIdFromNonAmazon();

        const callback = (e, url) => {
            Api.sendTracking(
                publisherHash,
                {
                    t: url.getAttribute('data-imp-tracking'),
                    ci_t: clId.type,
                    ci_v: clId.value,
                    r: Helpers.getReferer(),
                    a_u: window.location.href,
                    c_u: url.href,
                    u_agent: navigator.userAgent
                }
            );
        };

        Helpers.delegateEvent('click', callback, '.impact-track', document.body);
        ClIdTag._registerImpactMiddleClick(callback); 
    },
    handleNarrativTracking(publisherHash) {

        let clId = this._getClIdFromNonAmazon();
        const callback = (e, url) => {
            Api.sendTracking(
                publisherHash,
                {
                    t: url.getAttribute('data-nrv-tracking'),
                    ci_t: clId.type,
                    ci_v: clId.value,
                    r: Helpers.getReferer(),
                    a_u: window.location.href,
                    c_u: url.href,
                    u_agent: navigator.userAgent
                }
            );
        };

        Helpers.delegateEvent('click', callback, '.narrativ-track', document.body);
        ClIdTag._registerNarrativMiddleClick(callback);
    },
    _registerMiddleClick(callback) {
        let clicked = false;

        Helpers.delegateEvent('mouseup', (e, url) => {
            // Handle only a middle button.
            if (clicked) {
                callback.call(url, e, url);
                clicked = false;
            }
        }, '.carrick-track', document.body);
        Helpers.delegateEvent('mousedown', (e, url) => {
            // Handle only a middle button.
            if (e.button === 1) {
                clicked = true;
            }
        }, '.carrick-track', document.body);
    },
    _registerSkimlinksMiddleClick(callback) {
        let clicked = false;

        Helpers.delegateEvent('mouseup', (e, url) => {
            // Handle only a middle button.
            if (clicked) {
                callback.call(url, e, url);
                clicked = false;
            }
        }, '.skimlinks-track', document.body);
        Helpers.delegateEvent('mousedown', (e, url) => {
            // Handle only a middle button.
            if (e.button === 1) {
                clicked = true;
            }
        }, '.skimlinks-track', document.body);
    },
    _registerImpactMiddleClick(callback) {
        let clicked = false;

        Helpers.delegateEvent('mouseup', (e, url) => {
            // Handle only a middle button.
            if (clicked) {
                callback.call(url, e, url);
                clicked = false;
            }
        }, '.impact-track', document.body);
        Helpers.delegateEvent('mousedown', (e, url) => {
            // Handle only a middle button.
            if (e.button === 1) {
                clicked = true;
            }
        }, '.impact-track', document.body);
    },
    _registerNarrativMiddleClick(callback) {
        let clicked = false;

        Helpers.delegateEvent('mouseup', (e, url) => {
            // Handle only a middle button.
            if (clicked) {
                callback.call(url, e, url);
                clicked = false;
            }
        }, '.narrativ-track', document.body);
        Helpers.delegateEvent('mousedown', (e, url) => {
            // Handle only a middle button.
            if (e.button === 1) {
                clicked = true;
            }
        }, '.narrativ-track', document.body);
    },
    _rememberReferrer(force = false) {
        if (Cache.get('r', null) === null || force) {
            Cache.set('r', document.referrer, Config.cacheReferrerTtl)
        }
    },
    _getClIdFromUrl() {
        for (let clid in Config.clIds) {
            if (Helpers.checkUrlParameterExists(Config.clIds[clid])) {
                return {
                    type: Config.clIds[clid],
                    value: Helpers.getUrlParameter(Config.clIds[clid])
                };
            }
        }

        return null;
    },
    _getClIdFromNonAmazon(){
        let the_clid;
        let clIds  = Config.clIds;
        for (let clid in clIds) {
            if (Helpers.checkUrlParameterExists(clIds[clid])) {
                const cacheExpiry = Config[`cache${ClIdTag._tagType}Ttl`];
                const cachePath = Config[`cache${ClIdTag._tagType}Path`];

                const cacheValues = {
                    clid_t: clIds[clid],
                    clid_v: Helpers.getUrlParameter(clIds[clid])
                    // TODO: add subid
                };

                Cache.removeAll(Object.keys(cacheValues));
                for (let key in cacheValues) {
                    if (cacheValues[key]) {
                        Cache.set(key, cacheValues[key], cacheExpiry, cachePath);
                    }
                }
                ClIdTag._rememberReferrer(true);


                // Cache.set('clid_t', clIds[clid]);
                // Cache.set('clid_v', Helpers.getUrlParameter(clIds[clid]));

                the_clid = {
                    type: clIds[clid],
                    value: Helpers.getUrlParameter(clIds[clid])
                };
            }
        }
        if( !the_clid ){
            const clid_v = Cache.get('clid_v');

            if (clid_v) {
                the_clid = {
                    clid_t: Cache.get('clid_t'),
                    clid_v: clid_v
                }
            }
        }
        if (!the_clid) {
            the_clid = {
                type: null,
                value: null
            };
        }
        return the_clid;
    },
    _getClIdFromCache() {
        const clid_v = Cache.get('clid_v');

        if (clid_v) {
            return {
                clid_t: Cache.get('clid_t'),
                clid_v: clid_v,
                tag: Cache.get('tag'),
                r: Helpers.getReferer()
            }
        }

        return null
    },
    _createNewTag(clid) {
        const tag = Api.getTag(this._publisherHash);
        const cacheExpiry = Config[`cache${ClIdTag._tagType}Ttl`];
        const cachePath = Config[`cache${ClIdTag._tagType}Path`];

        return tag.then(result => {
            const cacheValues = {
                clid_t: clid.type,
                clid_v: clid.value,
                tag: result
            };

            Cache.removeAll(Object.keys(cacheValues));
            for (let key in cacheValues) {
                if (cacheValues[key]) {
                    Cache.set(key, cacheValues[key], cacheExpiry, cachePath);
                }
            }

            if (ClIdTag._tagType === 'Paid') {
                ClIdTag._rememberReferrer(true);
            }

            return result;
        });
    },
    _resolvePaid(resolve) {
        // get from url
        const clIdFromUrl = this._getClIdFromUrl();
        const clIdFromCache = this._getClIdFromCache();

        ClIdTag._tagType = 'Paid';

        if (clIdFromUrl && clIdFromCache) {
            if (clIdFromUrl.type == clIdFromCache.clid_t && clIdFromUrl.value == clIdFromCache.clid_v) {
                resolve(clIdFromCache.tag);
                return true;
            } else {
                resolve(this._createNewTag(clIdFromUrl))
                return true;
            }
        } else if (!clIdFromUrl && clIdFromCache) {
            resolve(clIdFromCache.tag)
            return true;
        } else if (clIdFromUrl && !clIdFromCache) {
            resolve(this._createNewTag(clIdFromUrl));
            return true;
        }

        return false;
    },
    _resolveOrganic(resolve) {
        const tag = Cache.get('tag');

        ClIdTag._tagType = 'Organic';

        if (!document.querySelector(Config.trackedSelector)) {
            resolve('');

            return;
        }

        if (!tag) {
            resolve(this._createNewTag({
                type: '',
                value: '',
                tag
            }));
        } else {
            resolve(tag);
        }

        return true;
    },
};

export {ClIdTag}