package main

/*
#cgo LDFLAGS: -lSDL3
#include <SDL3/SDL.h>

static inline SDL_Window* createWindow(const char* title, int w, int h) {
	return SDL_CreateWindow(title, w, h, 0);
}

static inline SDL_Renderer* createRenderer(SDL_Window* win) {
	SDL_Renderer* r = SDL_CreateRenderer(win, NULL);
	if (r) SDL_SetRenderVSync(r, 1);
	return r;
}

static inline void fillRect(SDL_Renderer* r, int x, int y, int w, int h) {
	SDL_FRect rect = {(float)x, (float)y, (float)w, (float)h};
	SDL_RenderFillRect(r, &rect);
}

static inline void drawRect(SDL_Renderer* r, int x, int y, int w, int h) {
	SDL_FRect rect = {(float)x, (float)y, (float)w, (float)h};
	SDL_RenderRect(r, &rect);
}

static inline void fillQuad(SDL_Renderer* r,
	float x0, float y0, float x1, float y1, float x2, float y2, float x3, float y3,
	float cr, float cg, float cb) {
	SDL_FColor col = {cr, cg, cb, 1.0f};
	SDL_Vertex v[4];
	v[0].position.x = x0; v[0].position.y = y0; v[0].color = col;
	v[1].position.x = x1; v[1].position.y = y1; v[1].color = col;
	v[2].position.x = x2; v[2].position.y = y2; v[2].color = col;
	v[3].position.x = x3; v[3].position.y = y3; v[3].color = col;
	int idx[6] = {0, 1, 2, 0, 2, 3};
	SDL_RenderGeometry(r, NULL, v, 4, idx, 6);
}

static inline int pollEvent(SDL_Event* e) {
	return SDL_PollEvent(e);
}

static inline Uint32 eventType(SDL_Event* e) {
	return e->type;
}

static inline SDL_Scancode eventKeyScancode(SDL_Event* e) {
	return e->key.scancode;
}

static inline int eventKeyRepeat(SDL_Event* e) {
	return e->key.repeat;
}

static inline int keyHeld(SDL_Scancode sc) {
	const bool* s = SDL_GetKeyboardState(NULL);
	return s[sc] ? 1 : 0;
}
*/
import "C"
import (
	"runtime"
	"unsafe"
)

func init() {
	runtime.LockOSThread()
}

type Window struct {
	ptr *C.SDL_Window
}

type Renderer struct {
	ptr *C.SDL_Renderer
}

type Event struct {
	raw C.SDL_Event
}

func SDLInit() bool {
	return bool(C.SDL_Init(C.SDL_INIT_VIDEO))
}

func SDLQuit() {
	C.SDL_Quit()
}

func SDLCreateWindow(title string, w, h int) *Window {
	ct := C.CString(title)
	defer C.SDL_free(unsafe.Pointer(ct))
	ptr := C.createWindow(ct, C.int(w), C.int(h))
	if ptr == nil {
		return nil
	}
	return &Window{ptr: ptr}
}

func SDLCreateRenderer(win *Window) *Renderer {
	ptr := C.createRenderer(win.ptr)
	if ptr == nil {
		return nil
	}
	return &Renderer{ptr: ptr}
}

func (r *Renderer) SetDrawColor(red, green, blue, alpha uint8) {
	C.SDL_SetRenderDrawColor(r.ptr, C.Uint8(red), C.Uint8(green), C.Uint8(blue), C.Uint8(alpha))
}

func (r *Renderer) Clear() {
	C.SDL_RenderClear(r.ptr)
}

func (r *Renderer) Present() {
	C.SDL_RenderPresent(r.ptr)
}

func (r *Renderer) FillRect(x, y, w, h int) {
	C.fillRect(r.ptr, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (r *Renderer) DrawRect(x, y, w, h int) {
	C.drawRect(r.ptr, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (r *Renderer) FillQuad(x0, y0, x1, y1, x2, y2, x3, y3 float64, red, green, blue uint8) {
	C.fillQuad(r.ptr,
		C.float(x0), C.float(y0), C.float(x1), C.float(y1),
		C.float(x2), C.float(y2), C.float(x3), C.float(y3),
		C.float(float64(red)/255), C.float(float64(green)/255), C.float(float64(blue)/255))
}

func (r *Renderer) Destroy() {
	C.SDL_DestroyRenderer(r.ptr)
}

func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.ptr)
}

func SDLPollEvent(e *Event) bool {
	return C.pollEvent(&e.raw) != 0
}

func (e *Event) Type() uint32 {
	return uint32(C.eventType(&e.raw))
}

func (e *Event) KeyScancode() int {
	return int(C.eventKeyScancode(&e.raw))
}

func (e *Event) KeyRepeat() bool {
	return C.eventKeyRepeat(&e.raw) != 0
}

func KeyHeld(scancode int) bool {
	return C.keyHeld(C.SDL_Scancode(scancode)) != 0
}

func SDLDelay(ms uint32) {
	C.SDL_Delay(C.Uint32(ms))
}

func SDLGetTicks() uint64 {
	return uint64(C.SDL_GetTicks())
}

const (
	SDL_EVENT_QUIT     = C.SDL_EVENT_QUIT
	SDL_EVENT_KEY_DOWN = C.SDL_EVENT_KEY_DOWN
)

const (
	SCANCODE_LEFT  = C.SDL_SCANCODE_LEFT
	SCANCODE_RIGHT = C.SDL_SCANCODE_RIGHT
	SCANCODE_DOWN  = C.SDL_SCANCODE_DOWN
	SCANCODE_UP    = C.SDL_SCANCODE_UP
	SCANCODE_SPACE = C.SDL_SCANCODE_SPACE
	SCANCODE_P     = C.SDL_SCANCODE_P
	SCANCODE_R     = C.SDL_SCANCODE_R
	SCANCODE_Q     = C.SDL_SCANCODE_Q
	SCANCODE_ESC   = C.SDL_SCANCODE_ESCAPE
)
