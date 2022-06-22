import { cloneDeep } from "lodash";
import { useStore } from "../store";
import ui from "../assets/wizard/installer/create-ui.json";
import model from "../assets/wizard/installer/model.json";
import schema from "../assets/wizard/installer/schema.json";
import language from "../assets/wizard/installer/language.json";
// @ts-ignore
import functions from "../assets/wizard/installer/functions.js";
import type {
  FunctionsInterface,
  LanguageInterface,
  ModelInterface,
  SchemaInterface,
  UiInterface,
} from "../types/ui-schemas";

export const uiBuilderJsonSetter = async () => {
  const store = useStore();
  let ui, model, schema, language;

  store.commit("wizard/ui$set", {});
  store.commit("wizard/schema$set", {});
  store.commit("wizard/model$init", {});
  store.commit("wizard/language$set", {});
  store.commit("wizard/functions$set", {});

  try {
    ui = await import("../assets/wizard/installer/create-ui.json");
    model = await import("../assets/wizard/installer/model.json");
    schema = await import("../assets/wizard/installer/schema.json");
    language = await import("../assets/wizard/installer/language.json");
  } catch (e) {
    console.log(e);
  }

  store.commit("wizard/ui$set", cloneDeep(ui.default));
  store.commit("wizard/schema$set", schema.default);
  store.commit("wizard/model$init", cloneDeep(model.default));
  store.commit("wizard/language$set", language.default);
  store.commit("wizard/functions$set", functions.default);
};
