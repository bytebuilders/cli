export function getServerDomain(
  subDomain?: string,
  inputPort?: string
): string {
  // execute only in client
  const hostname = window.location.hostname;
  const rootDomain = hostname.replace(
    new RegExp("^(console.|kubedb.|marketplace.|deploy.|grafana.|accounts.)"),
    ""
  );
  const protocol = window.location.protocol || "http";

  let host = `${protocol}//${subDomain ? subDomain + "." : ""}${rootDomain}`;
  if (hostname.search("bb.test") !== -1) {
    // dev
    const port = subDomain ? inputPort || window.location.port : "3000";
    host += port ? ":" + port : "";
  }

  return host;
}
