import {Api} from "./../api";
import {Helpers} from "./../helpers";

var subId1 = 'DG-' + Helpers.getAllUrlParams(document.querySelector("link[rel='shortlink']").href).p;

function ImpactTracker() {
	
	this.init = function () {

		this.handleVisit(process.env.MIX_PUBLISHER_ID)      
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

		var urls = document.querySelectorAll('a[href*="imp."]')
		for (let i = 0; i < urls.length; i++) {
			let url = urls[i];
			url.href = this.replaceSubIdParam(url.href);
			url.classList.add('impact-track');
		}
		this.handleTracking(process.env.MIX_PUBLISHER_ID)

	}
	
	this.replaceSubIdParam = function(urlString) {
		if (urlString.indexOf('subId1=') !== -1) {
			return urlString.replace(/([?&]subId1=)[^&]+/gi, '$1' + subId1);
		}
		let url = new URL(urlString);
		let search = url.search.substring(1);
		if (search) {
			search = '&' + search;
		}
		return url.origin + url.pathname + '?subId1=' + subId1 + '' + search;
	}

	this.handleTracking = function(publisherId) {
		// function get clid()
		const callback = (e, url) => {
			let clIdFromUrl = this.getClIdFromUrl();
			if (!clIdFromUrl) {
				clIdFromUrl = {
					type: null,
					value: null
				};
			}
			// TODO get the clid from cookies
			Api.sendTracking(publisherId, {
				t: subId1,
				ci_t: clIdFromUrl.type,
				ci_v: clIdFromUrl.value,
				r: document.referrer,
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

	this.handleVisit = function (publisherId) {
		// function get clid()
		let clIdFromUrl = this.getClIdFromUrl();
		if (!clIdFromUrl) {
			clIdFromUrl = {
				type: null,
				value: null
			};
		}
		Api.sendVisit(
			publisherId, 
			{
				t: subId1,
				ci_t: clIdFromUrl.type,
				ci_v: clIdFromUrl.value,
				r: document.referrer,
				a_u: window.location.href,
				u_agent: navigator.userAgent
			}
		);
	}
	this.getClIdFromUrl = function(){
		let clIds  = ['gclid', 'fbclid'];
		for (let clid in clIds) {
			if (Helpers.checkUrlParameterExists(clIds[clid])) {
				return {
					type: clIds[clid],
					value: Helpers.getUrlParameter(clIds[clid])
				};
			}
		}

		return null;
	}
}

const impactTracker = new ImpactTracker();

impactTracker.init();