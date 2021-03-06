package main

// DotTemplate ... Graph template for Graphviz
const DotTemplate = `
  digraph goDependencyGraph {
    label="{{ .Title }}";
    labeljust="1";
    fontname="Ubuntu";
    fontcolor="#222222";
    fontsize="13";
    bgcolor="lightsteelblue";
    style="solid";
    penwidth="1.0";
    pad="0.0";
    nodesep="1.0";

    node [shape="ellipse" style="filled" fillcolor="lightblue" fontname="Ubuntu" penwidth="1.0" margin="0.05, 0.0"];
    edge [color="#000000"]

    {{.DepsPath}}
  }
`