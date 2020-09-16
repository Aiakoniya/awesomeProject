package main

import (
	"fmt"
	"strings"
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}
type MarkdownListStrategy struct {}
func(m *MarkdownListStrategy) Start(builder *strings.Builder){}
func(m *MarkdownListStrategy) End(builder *strings.Builder){}
func(m *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string){
	builder.WriteString("*" + item + "\n")
}
type HtmlListStrategy struct{}
func(m *HtmlListStrategy) Start(builder *strings.Builder){
	builder.WriteString("<ul> \n")
}
func(m *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>")
}
func(m *HtmlListStrategy) AddListItem(builder *strings.Builder, item string){
	builder.WriteString(" <li>" + item +"</li>\n")
}
type TextProcessor struct{
	builder strings.Builder
	listStrategy ListStrategy
}
func(t *TextProcessor) SetOutputProcessor(strategy ListStrategy){
	t.listStrategy= strategy
}
func(t *TextProcessor) AppendList(items []string){
	t.listStrategy.Start(&t.builder)
	for _, item := range items {
		t.listStrategy.AddListItem(&t.builder, item)
	}
	t.listStrategy.End(&t.builder)
}
func(t *TextProcessor) Reset() {
	t.builder.Reset()
}
func(t *TextProcessor) String() string {
	return t.builder.String()
}
func NewTextProcessor(strategy ListStrategy) *TextProcessor {
	return &TextProcessor{strings.Builder{}, strategy}
}
func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"AITU","KBTU", "IITU"})
	fmt.Println(tp)
	tp.Reset()
	tp.SetOutputProcessor(&HtmlListStrategy{})
	tp.AppendList([] string{"AITU", "KBTU", "IITU"})
	fmt.Println(tp)
}