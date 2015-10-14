package main

type stacking int

const (
	Above stacking = iota
	Below
)

type stack struct {
	n    *Node
	prev *stack
	next *stack
}

func newStack(n *Node) *stack {
	return &stack{n, nil, nil}
}

//func (s *stack) pop(Direction) *stack {}

func (s *stack) push(Direction) {}

/*
void stack_insert_after(stacking_list_t *a, node_t *n)
{
	stacking_list_t *s = make_stack(n);
	if (a == NULL) {
		stack_head = stack_tail = s;
	} else {
		if (a->node == n)
			return;
		remove_stack_node(n);
		stacking_list_t *b = a->next;
		if (b != NULL)
			b->prev = s;
		s->next = b;
		s->prev = a;
		a->next = s;
		if (stack_tail == a)
			stack_tail = s;
	}
}

void stack_insert_before(stacking_list_t *a, node_t *n)
{
	stacking_list_t *s = make_stack(n);
	if (a == NULL) {
		stack_head = stack_tail = s;
	} else {
		if (a->node == n)
			return;
		remove_stack_node(n);
		stacking_list_t *b = a->prev;
		if (b != NULL)
			b->next = s;
		s->prev = b;
		s->next = a;
		a->prev = s;
		if (stack_head == a)
			stack_head = s;
	}
}

void remove_stack(stacking_list_t *s)
{
	if (s == NULL)
		return;
	stacking_list_t *a = s->prev;
	stacking_list_t *b = s->next;
	if (a != NULL)
		a->next = b;
	if (b != NULL)
		b->prev = a;
	if (s == stack_head)
		stack_head = b;
	if (s == stack_tail)
		stack_tail = a;
	free(s);
}

void remove_stack_node(node_t *n)
{
	for (stacking_list_t *s = stack_head; s != NULL; s = s->next)
		if (s->node == n) {
			remove_stack(s);
			return;
		}
}

void stack(node_t *n, stack_flavor_t f)
{
	PRINTF("stack %X\n", n->client->window);

	if (stack_head == NULL) {
		stack_insert_after(NULL, n);
	} else if (n->client->fullscreen) {
		if (f == STACK_ABOVE) {
			stack_insert_after(stack_tail, n);
			window_raise(n->client->window);
		}
	} else {
		if (f == STACK_ABOVE && n->client->floating && !auto_raise)
			return;
		stacking_list_t *latest_tiled = NULL;
		stacking_list_t *oldest_floating = NULL;
		stacking_list_t *oldest_fullscreen = NULL;
		for (stacking_list_t *s = (f == STACK_ABOVE ? stack_tail : stack_head); s != NULL; s = (f == STACK_ABOVE ? s->prev : s->next)) {
			if (s->node != n) {
				if (s->node->client->fullscreen) {
					if (oldest_fullscreen == NULL)
						oldest_fullscreen = s;
					continue;
				}
				if (s->node->client->floating == n->client->floating) {
					if (f == STACK_ABOVE) {
						stack_insert_after(s, n);
						window_above(n->client->window, s->node->client->window);
					} else {
						stack_insert_before(s, n);
						window_below(n->client->window, s->node->client->window);
					}
					return;
				} else if ((f != STACK_ABOVE || latest_tiled == NULL) && !s->node->client->floating) {
					latest_tiled = s;
				} else if ((f == STACK_ABOVE || oldest_floating == NULL) && s->node->client->floating) {
					oldest_floating = s;
				}
			}
		}
		if (latest_tiled == NULL && oldest_floating == NULL && oldest_fullscreen == NULL)
			return;
		if (n->client->floating) {
			if (latest_tiled != NULL) {
				window_above(n->client->window, latest_tiled->node->client->window);
				stack_insert_after(latest_tiled, n);
			} else if (oldest_fullscreen != NULL) {
				window_below(n->client->window, oldest_fullscreen->node->client->window);
				stack_insert_before(oldest_fullscreen, n);
			}
		} else {
			if (oldest_floating != NULL) {
				window_below(n->client->window, oldest_floating->node->client->window);
				stack_insert_before(oldest_floating, n);
			} else if (oldest_fullscreen != NULL) {
				window_below(n->client->window, oldest_fullscreen->node->client->window);
				stack_insert_before(oldest_fullscreen, n);
			}
		}
	}
}
*/
