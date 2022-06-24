import type { InjectionKey } from "vue";
import { createStore, useStore as baseUseStore, Store } from "vuex";

export interface State {}

export const key: InjectionKey<Store<State>> = Symbol();

export const store: Store<State> = createStore<State>({
  state: {},

  mutations: {},

  actions: {},

  //plugins: process.env.NODE_ENV !== "production" ? [createLogger()] : [],
  //strict: process.env.NODE_ENV !== "production",
});

// define custom `useStore` composition function
export function useStore(): Store<State> {
  return baseUseStore(key);
}
