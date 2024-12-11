const Api = {
    baseUrl: process.env.MIX_API_URL,
	sendVisit: function sendVisit(publisherHash, _ref) {
		var t = _ref.t,
		ci_t = _ref.ci_t,
		ci_v = _ref.ci_v,
		r = _ref.r,
		a_u = _ref.a_u,
		u_agent = _ref.u_agent;
		var data = {
			t: t,
			ci_t: ci_t,
			ci_v: ci_v,
			r: r,
			a_u: a_u,
			u_agent: u_agent
		};
		return Api.sendApi("".concat(Api.baseUrl, "/visit/").concat(publisherHash), data);
	},
	sendTracking: function sendTracking(publisherHash, _ref2) {
		var t = _ref2.t,
		ci_t = _ref2.ci_t,
		ci_v = _ref2.ci_v,
		r = _ref2.r,
		a_u = _ref2.a_u,
		c_u = _ref2.c_u,
		u_agent = _ref2.u_agent;
		var data = {
			t: t,
			ci_t: ci_t,
			ci_v: ci_v,
			r: r,
			a_u: a_u,
			c_u: c_u,
			u_agent: u_agent
		};
		return Api.sendApi("".concat(Api.baseUrl, "/send-tracking/").concat(publisherHash), data);
	},
  // FIXME: In the possible future.
  //  This method can return an unexpected value depending on the browser.
  sendApi: function sendApi(url, data) {
  	if (navigator.sendBeacon) {
  		return navigator.sendBeacon(url, JSON.stringify(data));
  	}

  	var req = new XMLHttpRequest();
  	req.open('POST', url, false);
  	req.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
  	req.send(JSON.stringify(data));
  	return req.status === 200;
  }
};
export {Api}