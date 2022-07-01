<template>
  <div class="bb-installer mt-70">
    <root-navbar />
    <div class="container-720">
      <div class="ac-markdown">
        <div v-html="mdToHtml"></div>
        <div
          class="is-flex is-justify-content-space-between is-align-items-center"
        >
          <button
            class="button ac-button is-small is-secondary"
            @click="onBackClick"
          >
            <i class="fa fa-arrow-left mr-15"></i> Back
          </button>
          <button
            class="button ac-button is-small is-primary"
            @click="onDownloadClick"
          >
            Downloads <i class="fa fa-download ml-15"></i>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import RootNavbar from "./RootNavbar.vue";
import { computed } from "vue";
import { useStore } from "./../store";
import { useRouter } from "vue-router";
import { marked } from "marked";

//init store & router
const store = useStore();
const router = useRouter();

// Get markdown from the store
const markDown = computed(() => store.getters["getMarkDown"]);

const onBackClick = () => {
  router.push("./");
};

const onDownloadClick = () => {
  window.print();
};

const mdToHtml = marked.parse(markDown.value);
</script>

<style lang="scss">
.container-720 {
  width: 720px;
  margin: 0 auto;
}

.ac-markdown {
  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    line-height: 1.3;
    font-weight: 500;
    margin-bottom: 15px;
    border-bottom: 1px solid #e7e7e7;
    padding-bottom: 10px;
  }
  h1 {
    font-size: 28px;
  }
  h2 {
    font-size: 24px;
  }
  h3 {
    font-size: 20px;
  }
  h4 {
    font-size: 18px;
  }

  h5 {
    font-size: 16px;
  }
  h5 {
    font-size: 14px;
  }
  p {
    margin-bottom: 30px;
    color: #313131;
    font-size: 16px;
    strong {
      font-weight: 400;
      color: #313131;
      background: #f6f8fa;
      padding: 15px;
      display: inline-block;
      border-radius: 4px;
    }
  }

  ul {
    list-style: disc;
    margin-left: 30px;
    padding: 0 0 0 30px;
    margin-bottom: 30px;
    li {
      display: block;
      font-size: 16px;
      color: #313131;
      padding-bottom: 8px;
      text-indent: -30px;
      &:before {
        content: "\f10c";
        font-family: fontawesome;
        font-size: 8px;
        margin-right: 20px;
        color: #313131;
      }
    }
  }
  table {
    margin-bottom: 30px;
    tr {
      border-left: 1px solid #e7e7e7;
      border-top: 1px solid #e7e7e7;
      &:nth-child(even) {
        background-color: #ddd;
      }
      th {
        font-weight: 500;
      }
      td,
      th {
        border-right: 1px solid #e7e7e7;
        border-bottom: 1px solid #e7e7e7;
        padding: 10px 20px;
        font-size: 14px;
        a {
          color: #1971bd;
          font-weight: 500;
        }
      }
    }
  }

  pre {
    margin-bottom: 15px;
    code {
      font-size: 14px;
    }
  }
}
</style>
