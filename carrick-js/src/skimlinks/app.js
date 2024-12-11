function SkimlinksTracker() {
	
	this.init = function () {
		if (document.readyState !== "loading") {
			console.log("loading");
			this.initHandle();
		}
		window.addEventListener('DOMContentLoaded', () => {
			console.log("DOMContentLoaded");
			this.initHandle();
		});
		window.addEventListener('load', () => {
			console.log("load");
			this.initHandle()
		});
	};

	this.initHandle = function () {
		console.log( document.getElementById("skimlinks-pixels-iframe") );
		if ( document.getElementById("skimlinks-pixels-iframe") ){
			// console.log('skimlinks is installed !')
			this.amazonSelector = 'a[href*="amazon."]';
			this.nonAmazonSelectors = 'a[href*="ebay."]';
			this.transformAmazonUrls();
			this.transformNonAmazonrUrls();
		}
	};

	this.transformAmazonUrls = function() {
		let urls = document.querySelectorAll(this.amazonSelector)
		for (let i = 0; i < urls.length; i++) {
			let url = urls[i];
			url.classList.add('noskim');
		}
	}
	this.transformNonAmazonrUrls = function(){
		let urls = document.querySelectorAll(this.nonAmazonSelectors)
		for (let i = 0; i < urls.length; i++) {
			let url = urls[i];
			let custom_id = 'DG-' + this.hashPath(new URL(url.href).pathname)
			url.setAttribute('data-skimlinks-tracking', custom_id)
		}

	}
	this.hashPath = function(path){
		return path.split("").reduce(function(a,b){a=((a<<5)-a)+b.charCodeAt(0);return a&a},0);              
	}

	this.getParamValue = function(link, param) {
		return new URL(link).searchParams.get(param);
	}

}

const skimlinksTracker = new SkimlinksTracker();

skimlinksTracker.init();