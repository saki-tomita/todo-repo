// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import { createVuetify } from "vuetify";
import * as labs from 'vuetify/labs/components'
//import * as components from "vuetify/components";
// import * as directives from "vuetify/directives";
// import { aliases, mdi } from "vuetify/iconsets/mdi";
// import { VDataTable } from 'vuetify/labs/VDataTable';

// const vuetify = createVuetify({
//     components: {
//         VDataTable
//     },
//     directives,
//     icons: {
//         defaultSet: "mdi",
//         aliases,
//         sets: {
//             mdi,
//         },
//     },
// });
//
// export default vuetify;


// VDataTableのインポートを新規追加
// import { VDataTable } from 'vuetify/labs/VDataTable'



// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
    // componentsを新規追加
    components: {
        // VDataTable,
        ...labs,
    },
    theme: {
        themes: {
            light: {
                colors: {
                    primary: '#1867C0',
                    secondary: '#5CBBF6',
                },
            },
        },
    },

})

