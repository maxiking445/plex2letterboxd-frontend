import { createApp } from "vue";
import App from "./App.vue";
import "../assets/main.css";
import "vue-toastification/dist/index.css";
import Toast from "vue-toastification";

// FontAwesome
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

// Icons importieren
import { faPlay, faGear } from "@fortawesome/free-solid-svg-icons";

library.add(faPlay, faGear);

createApp(App)
  .component("font-awesome-icon", FontAwesomeIcon)
  .use(Toast, {
    position: "top-center",
    timeout: 3000,
    pauseOnHover: true,
  })
  .mount("#app");
