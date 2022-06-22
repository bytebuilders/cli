import { cloneDeep } from "lodash";
import { useStore } from "../store";

import ui from "../assets/wizard/installer/create-ui.yaml";
import model from "../assets/wizard/installer/model.yaml";
import schema from "../assets/wizard/installer/schema.yaml";
import language from "../assets/wizard/installer/language.yaml";
// @ts-ignore../assets/wizard/installer/language.yaml
import functions from "../assets/wizard/installer/functions.js";

export const uiBuilderJsonSetter = async () => {
  const store = useStore();

  store.commit("wizard/ui$set", cloneDeep(ui));
  store.commit("wizard/schema$set", schema);
  store.commit("wizard/model$init", cloneDeep(model));
  store.commit("wizard/language$set", language);
  store.commit("wizard/functions$set", functions);
};
