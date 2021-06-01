import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes: {
            light: {
                error: '#ff3838',
                success: '#3aa500',
                info: '#00e7d7',
                warning: '#957903',
                primary: '#3f2bfa',
                secondary: '#1f1f1f',
                accent: '#82B1FF',
            },
        },
    },
});
