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
    const uiFile = await import(`../${name}/ui/create-ui.yaml`);
    ui = uiFile.default || {};
    const languageFile = await import(`../${name}/ui/language.yaml`);
    language = languageFile.default || {};
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
