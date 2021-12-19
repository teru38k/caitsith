import colors from 'vuetify/es5/util/colors'

export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
      title: 'WebApp',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    {
      src: '@/plugins/imageselect',
      mode: 'client'
    },
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    // https://go.nuxtjs.dev/vuetify
    '@nuxtjs/vuetify'
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    '@nuxtjs/proxy',
    // https://go.nuxtjs.dev/pwa
    '@nuxtjs/pwa'
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    proxy: true,
  },
   
  proxy: {
    '/api': {
      target: 'http://host.docker.internal:1323',
      pathRewrite: {
        '^/api': ''
      },
    }
  },

  // PWA module configuration: https://go.nuxtjs.dev/pwa
  pwa: {
    manifest: {
      lang: 'en'
    }
  },

  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    treeShake: true,
    theme: {
      options: {
        customProperties: true
      },
      light: true,
      themes: {
        light: {
          background: '#00acee',
          primary: '#00ced1',
          secondary: '#f08080',
          accent: '#9370db',
          error: '#2f4f4f', 
          info: '#008080',
          warning: '#b22222',
          success: '#7cfc00',
        }
      },
      //dark: false,
      //themes: {
      //  dark: {
      //    primary: colors.blue.darken2,
      //    accent: colors.grey.darken3,
      //    secondary: colors.amber.darken3,
      //    info: colors.teal.lighten1,
      //    warning: colors.amber.base,
      //    error: colors.deepOrange.accent4,
      //    success: colors.green.accent3
      //  }
      //}
    }
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  }
}
