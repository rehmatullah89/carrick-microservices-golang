import {Helpers} from "./helpers";
import {ClIdTag} from "./clid";
import {Config} from "./config";

function CarricTracker() {
    this.tag = '';
    this.skimlinks_activated = false;
    this.tagInProgress = false;

    this.init = function () {
        ClIdTag.handleVisit(process.env.MIX_PUBLISHER_ID);

        this.initObserver(); // Observe href attribute for changes
        if (document.readyState !== "loading") {
            this.initHandle();
        }
        window.addEventListener('DOMContentLoaded', () => {
            this.initHandle();
        });
        window.addEventListener('load', () => {
            this.initSkimlinksHandle();
            this.initImpactHandle();
            this.initNarrativHandle();
            this.initHandle();
        });
    };

    this.initHandle = function () {
        // TODO 
        // Move this code below to this.handle
        if (this.tagInProgress) {
            return;
        }
        if (this.tag) {
            Helpers.transformUrls(this.tag);
            return;
        }

        this.handle();
    };

    this.handle = function () {
        this.tagInProgress = true;

        // if url has gclid, fbclid, msclkid
        const publisherId = process.env.MIX_PUBLISHER_ID;
        const tagPromise = ClIdTag.get(publisherId);

        tagPromise.then(tag => {
            // replace tags in links
            if (tag) {
                this.tag = tag;
                Helpers.transformUrls(this.tag);
                ClIdTag.handleTracking(publisherId);
            }

            this.tagInProgress = false;
        });
    }

    this.initObserver = function () {
        let maxMutations = Config.observerMaxMutations;
        const regex = new RegExp(Config.trackedUrls[0]);
        const observer = new MutationObserver((records, observer) => {
            let record;
            if (maxMutations <= 0) {
                observer.disconnect();
                return;
            }

            maxMutations--;
            if (!this.tag) {
                return;
            }

            for (let i = 0; i < records.length; i++) {
                const tagRegex = new RegExp('[?\&]tag\=' + this.tag);

                record = records[i];
                if (record.type !== 'attributes' ||
                    record.target.tagName !== 'A' ||
                    !record.target.href) {
                    continue;
                }

                const href = record.target.href;

                if (regex.test(href) && !tagRegex.test(href)) {
                    Helpers.transformUrls(this.tag, [record.target]);
                }
            }
        });

        observer.observe(document.body, {
            attributes: true,
            attributeFilter: ['href'],
            subtree: true,
        });
    }
    this.initImpactHandle = function(){
        const publisherId = process.env.MIX_PUBLISHER_ID;
        Helpers.transformImpactUrls();
        ClIdTag.handleImpactTracking(publisherId);
    }
    this.initSkimlinksHandle = function() {
        const publisherId = process.env.MIX_PUBLISHER_ID;
        Helpers.transformSkimlinksUrls();
        ClIdTag.handleSkimlinksTracking(publisherId);
    }
    this.initNarrativHandle = function() {
        const publisherId = process.env.MIX_PUBLISHER_ID;
        Helpers.transformNarrativUrls();
        ClIdTag.handleNarrativTracking(publisherId);
    }
}


const carrickTracker = new CarricTracker();

carrickTracker.init();