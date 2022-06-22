import { createI18n } from "vue-i18n";
import en from "../locales/en.json";
import bn from "../locales/bn.json";

const i18n = createI18n({
  locale: "en",
  fallbackLocale: "en",
  messages: { en, bn },
});

export default i18n;
