package ichido

type ContentHolder struct{
    Value interface{}
}

type ContentsProvider func() []*ContentHolder
type NewlyChecker func(content *ContentHolder) bool
type Invoker func(content *ContentHolder)
type InvokedMarker func(content *ContentHolder) 

type Ichido struct{
    contentProvider ContentsProvider
    newlyChecker NewlyChecker
    invoker Invoker
    invokedMarker InvokedMarker 
}
 
func (i *Ichido) RegisterContentsProvider(provider ContentsProvider) {
    i.contentProvider = provider 
}

func (i *Ichido) RegisterNewlyChecker(checker NewlyChecker) {
    i.newlyChecker = checker
}

func (i *Ichido) RegisterInvoker(invoker Invoker) {
    i.invoker = invoker
}

func (i *Ichido) RegisterInvokedMarker(marker InvokedMarker) {
    i.invokedMarker = marker
}

func (i *Ichido) Run() {
    for _, content := range i.contentProvider() {
        if !i.newlyChecker(content) {
            continue
        }

        i.invoker(content)
        i.invokedMarker(content)
    }    
}