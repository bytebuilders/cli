import $axios from "@/plugins/axios";
import { cloneDeep } from "lodash";
import { useStore } from "../store";

export const uiBuilderJsonSetter = async () => {
  //init vuex
  const store = useStore();

  //declare interface
  interface uiType {
    default: Record<string, unknown>;
  }

  //declare variable
  let ui: uiType = { default: {} };
  let schema: uiType = { default: {} };
  let model: uiType = { default: {} };
  let functions = {};
  let language: uiType = { default: {} };

  //clear ui-builder store
  store.commit("wizard/ui$set", {});
  store.commit("wizard/schema$set", {});
  store.commit("wizard/model$init", {});
  store.commit("wizard/language$set", {});
  store.commit("wizard/functions$set", {});

  try {
    //get schema & model from api call
    const schemaReq = await $axios.get("/apis/schema.json");
    schema = schemaReq.data || {};
    const modelReq = await $axios.get("/apis/options.json");
    model = modelReq.data || {};

    //read yaml from file
    ui = await import("../assets/wizard/installer/create-ui" + ".yaml");
    language = await import("../assets/wizard/installer/language" + ".yaml");
    functions = await import("../assets/wizard/installer/functions" + ".js");
  } catch (error) {
    console.log(error);
  }

  //set value to vuex
  store.commit("wizard/schema$set", schema);
  store.commit("wizard/model$init", model);

  store.commit("wizard/ui$set", cloneDeep(ui.default));
  store.commit("wizard/language$set", language.default);
  store.commit("wizard/functions$set", functions);
};
