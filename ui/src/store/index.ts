import type { InjectionKey } from "vue";
import { createStore, useStore as baseUseStore, Store } from "vuex";

export interface State {
  markDown: string;
  isDataAvailable: boolean;
}

export const key: InjectionKey<Store<State>> = Symbol();

export const store: Store<State> = createStore<State>({
  state: {
    markDown: "",
    isDataAvailable: false,
  },

  getters: {
    getMarkDown(state) {
      return state.markDown || "";
    },
    getIsDataAvailable(state) {
      return state.isDataAvailable || false;
    },
  },

  mutations: {
    setMarkDown(state, value) {
      state.markDown = value || {};
    },
    setIsDataAvailable(state, value: boolean) {
      state.isDataAvailable = value;
    },
  },

  actions: {},

  //plugins: process.env.NODE_ENV !== "production" ? [createLogger()] : [],
  //strict: process.env.NODE_ENV !== "production",
});

// define custom `useStore` composition function
export function useStore(): Store<State> {
  return baseUseStore(key);
}
