async function loadLocalFile(url) {
  try {
    const finalUrl = `../${url}`;
    const localFile = await import(finalUrl);
    return localFile.default;
  } catch (e) {
    return e;
  }
}
async function loadFunctionJs(name) {
  try {
    const globFunctionJs = await import.meta.globEager("../*/ui/functions.js");
    const functionjs = globFunctionJs["../" + name + "/ui/functions.js"];
    return functionjs;
  } catch (e) {
    console.log(e);
    throw e;
  }
}
export async function loadLocalComponent({ itemCtx }) {
  const name = `${itemCtx.chart.name}`;

  let ui = {};
  let language = {};
  let functions = {};

  try {
    ui = await loadLocalFile(`${name}/ui/create-ui.yaml`);
    language = await loadLocalFile(`${name}/ui/language.yaml`);
    functions = await loadFunctionJs(name);
  } catch (e) {
    console.log(e);
  }

  return {
    ui,
    language,
    functions,
  };
}
