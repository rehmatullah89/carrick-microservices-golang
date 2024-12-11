# Environment configuration
`cp .env.example .env`

In the .env file you need to specify your publisher id \
`MIX_PUBLISHER_ID=your_publisher_id`

# How to build

---

``npm run production`` - for production

``npm run dev`` - for development

``npm run watch`` - to watch and compile files once they have changed

The output file will appear in `dist/` directory.

# Running examples

---

``npm run hot`` - runs server with examples in `examples/` directory.

# Widget

---
````
<script src="/tracking.js" async="true"></script>
<script type="text/javascript">
    (function(w) {
        w.advonCommerce = w.advonCommerce || [];
        function init(){advonCommerce.push(arguments);}
        init('tracker', {
            publisher: 'c4a675af07781d21b0488cb1fc23cea1'
        });
    })(window)
</script>
````