import {Api} from "./../api";
import {Helpers} from "./../helpers";
import {Cache} from "./../cache";

function ImpactTracker() {
	
	this.init = function () {

		if (document.readyState !== "loading") {
			this.initHandle();
		}
		window.addEventListener('DOMContentLoaded', () => {
			this.initHandle();
		});
		window.addEventListener('load', () => this.initHandle());
	};

	this.initHandle = function () {
		this.transformUrls();
	};

	this.transformUrls = function() {

		var urls = document.querySelectorAll(Helpers.listImpact)
		for (let i = 0; i < urls.length; i++) {
			let url = urls[i];
			// url.href = this.replaceSubIdParam(url);
			url.classList.add('impact-track');
			// skip skimlinks redir
			// if ( document.getElementById("skimlinks-pixels-iframe") ){
			// 	url.classList.add('noskim');
			// }
		}
		this.handleTracking(process.env.MIX_PUBLISHER_ID)

	}
	
	this.replaceSubIdParam = function(url) {
		let subId1 = 'DG' + this.hashPath(url.pathname);

		if (url.href.indexOf('subId1=') !== -1) {
			return url.href.replace(/([?&]subId1=)[^&]+/gi, '$1' + subId1);
		}
		// let url = new URL(urlString);
		let search = url.search.substring(1);
		if (search) {
			search = '&' + search;
		}
		return url.origin + url.pathname + '?subId1=' + subId1 + '' + search;
	}

	this.hashPath = function(path){
		return path.split("").reduce(function(a,b){a=((a<<5)-a)+b.charCodeAt(0);return a&a},0);              
	}

	this.handleTracking = function(publisherId) {
		// function get clid()
		const callback = (e, url) => {
			
			let subId1 = 'DG' + this.hashPath(url.pathname);
        	
			let clId = this.getClIdFromUrl();
        	if( !clId ){
        		clId = this.getClIdFromCache();
        	}

			if (!clId) {
				clId = {
					type: null,
					value: null
				};
			}

			// TODO get the clid from cookies
			Api.sendTracking(publisherId, {
				t: subId1,
				ci_t: clId.type,
				ci_v: clId.value,
				r: this.getReferer(),
				a_u: window.location.href,
				c_u: url.href,
				u_agent: navigator.userAgent
			});
		};

		this.delegateEvent('click', callback, '.impact-track', document.body);
		this._registerMiddleClick(callback);
	}
	
	this.delegateEvent = function delegateEvent(event, callback, targetSelector, parentElement) {
		parentElement.addEventListener(event, function (e) {
			var elem = e.target.closest(targetSelector);

			if (elem && callback instanceof Function) {
				callback.call(elem, e, elem);
			}
		});
	}
	this._registerMiddleClick = function _registerMiddleClick(callback) {
		var clicked = false;
		
		this.delegateEvent('mouseup', function (e, url) {
			// Handle only a middle button.
			if (clicked) {
				callback.call(url, e, url);
				clicked = false;
			}
		}, '.impact-track', document.body);
		
		this.delegateEvent('mousedown', function (e, url) {
			// Handle only a middle button.
			if (e.button === 1) {
				clicked = true;
			}
		}, '.impact-track', document.body);
	}

	this.getClIdFromUrl = function(){
		let clIds  = ['gclid', 'fbclid'];
		for (let clid in clIds) {
			if (Helpers.checkUrlParameterExists(clIds[clid])) {
				Cache.set('clid_t', clIds[clid]);
				Cache.set('clid_v', Helpers.getUrlParameter(clIds[clid]));

				return {
					type: clIds[clid],
					value: Helpers.getUrlParameter(clIds[clid])
				};
			}
		}

		return null;
	}
	this.getClIdFromCache = function() {
        const clid_v = Cache.get('clid_v');

        if (clid_v) {
            return {
                type: Cache.get('clid_t'),
                value: clid_v,
            }
        }

        return null
    }
    this.getReferer = function(){
    	return document.referrer ? document.referrer : window.location.origin;
    }
}

const impactTracker = new ImpactTracker();

impactTracker.init();