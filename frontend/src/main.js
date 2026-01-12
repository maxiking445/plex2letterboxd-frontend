import { createApp } from "vue";
import App from "./App.vue";
import "../assets/main.css";
import "vue-toastification/dist/index.css";
import Toast from "vue-toastification";
import { VueSpinnersPlugin } from "vue3-spinner";

// FontAwesome
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

// Icons import
import {
  faPlay,
  faGear,
  faVial,
  faSave,
  faDownload,
  faCheck,
  faTrash,
  faEyeSlash,
  faEye,
} from "@fortawesome/free-solid-svg-icons";

library.add(
  faPlay,
  faGear,
  faVial,
  faSave,
  faDownload,
  faCheck,
  faTrash,
  faEyeSlash,
  faEye
);

createApp(App)
  .use(VueSpinnersPlugin)
  .component("font-awesome-icon", FontAwesomeIcon)
  .use(Toast, {
    position: "top-center",
    timeout: 3000,
    pauseOnHover: true,
  })
  .mount("#app");
