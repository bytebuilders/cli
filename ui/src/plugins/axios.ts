import axios from "axios";
import type { AxiosInstance } from "axios";
import { getServerDomain } from "./domain";
import { get as getCookie } from "tiny-cookie";

const $axios: AxiosInstance = axios.create({
  baseURL: `/`,
  timeout: 1000000,
  headers: {
    "x-csrf-token": getCookie("_csrf"),
    "X-Requested-With": "xmlhttprequest",
  },
  withCredentials: true,
});

export default $axios;
