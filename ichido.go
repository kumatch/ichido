package ichido

type ContentsProvider func() []interface{}
type NewlyChecker func(content interface{}) bool
type Invoker func(content interface{})
type InvokedMarker func(content interface{})

type Ichido struct {
	contentProvider ContentsProvider
	newlyChecker    NewlyChecker
	invoker         Invoker
	invokedMarker   InvokedMarker
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
