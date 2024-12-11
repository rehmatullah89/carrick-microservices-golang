const Cache = {
    set(key, value, ttl = 60, path = '/') {
        document.cookie = `${Cache._getKey(key)}=${value};path=${path};max-age=${ttl};`;
    },
    get(key, def = '') {
        const cookie = Cache._parseCookie(key);

        if (!cookie) {
            return def;
        }

        return cookie.split('=')[1];
    },
    removeAll(keys, path) {
        if (!(keys instanceof Array)) {
            keys = [keys];
        }

        keys.forEach((key) => {
            if (path) {
                Cache.remove(key, path);

                return;
            }

            Cache.remove(key, '/');
            Cache.remove(key, document.location.pathname);
        });
    },
    remove(key, path = '/') {
        document.cookie = `${Cache._getKey(key)}=;max-age=0;path=${path}`;
    },
    _getKey(key) {
        return `${process.env.MIX_CACHE_KEY_PREFIX}_${key}`;
    },
    _parseCookie(key) {
        return document.cookie.split('; ')
            .find(row => row.startsWith(`${Cache._getKey(key)}=`));
    }
};

export {Cache}