function HideTopbarPlugin() {
  // this plugin overrides the Topbar component to return nothing
  return {
    components: {
      Topbar: function() { return null }
    }
  }
}

window.onload = function() {

  // Build a system
  const ui = SwaggerUIBundle({
    url: "https://mosmix-api.codeformuenster.org/mosmix-api.json",
    dom_id: '#swagger-ui',
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl,
      HideTopbarPlugin
    ],
    layout: "StandaloneLayout"
  })

  window.ui = ui
}
