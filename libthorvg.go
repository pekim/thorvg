package thorvg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/ebitengine/purego"
)

func libraryFilepath() (string, error) {
	filepath := filepath.Join(os.TempDir(), fmt.Sprintf("libthorvg-%s.so", sharedObjectID))

	// Check if the file exists. If it does, don't create it.
	//
	// The code would be simpler, and very nearly as quick, if the check were omitted and the
	// file written every time. However that causes problems if multiple applications are run
	// concurrently. The replaced file becomes unavailable, and a segment violation results.
	_, err := os.Stat(filepath)
	if err == nil {
		return filepath, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("failed to check existence of shared object file %q : %w", filepath, err)
	}

	err = os.WriteFile(filepath, sharedObject, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to write shared object file %q : %w", filepath, err)
	}

	return filepath, nil
}

// Engine API
var tvg_engine_init func(threads int) Result
var tvg_engine_term func() Result
var tvg_engine_version func(major *uint32, minor *uint32, micro *uint32, version **byte) Result

// Canvas API
var tvg_swcanvas_create func(option EngineOption) uintptr
var tvg_swcanvas_set_target func(canvas uintptr, buffer *byte, stride uint32, w uint32, h uint32, cs ColorSpace) Result
var tvg_glcanvas_create func(option EngineOption) uintptr
var tvg_glcanvas_set_target func(canvas uintptr, display unsafe.Pointer, surface unsafe.Pointer, context unsafe.Pointer, id int32, w uint32, h uint32, cs ColorSpace) Result
var tvg_wgcanvas_create func() uintptr
var tvg_wgcanvas_set_target func(canvas uintptr, device unsafe.Pointer, instance unsafe.Pointer, target unsafe.Pointer, w uint32, h uint32, cs ColorSpace, typ int32) Result
var tvg_canvas_destroy func(canvas uintptr) Result
var tvg_canvas_add func(canvas uintptr, paint uintptr) Result
var tvg_canvas_insert func(canvas uintptr, target uintptr, paint uintptr) Result
var tvg_canvas_remove func(canvas uintptr, paint uintptr) Result
var tvg_canvas_update func(canvas uintptr) Result
var tvg_canvas_draw func(canvas uintptr, clear_ bool) Result
var tvg_canvas_sync func(canvas uintptr) Result
var tvg_canvas_set_viewport func(canvas uintptr, x int32, y int32, w int32, h int32) Result

// Paint API
var tvg_paint_rel func(paint uintptr) Result
var tvg_paint_ref func(paint uintptr) uint16
var tvg_paint_unref func(paint uintptr, free bool) uint16
var tvg_paint_get_ref func(paint uintptr) uint16
var tvg_paint_set_visible func(paint uintptr, visible bool) Result
var tvg_paint_get_visible func(paint uintptr) bool
var tvg_paint_get_id func(paint uintptr) uint32
var tvg_paint_set_id func(paint uintptr, id uint32) Result
var tvg_paint_scale func(paint uintptr, factor float32) Result
var tvg_paint_rotate func(paint uintptr, degree float32) Result
var tvg_paint_translate func(paint uintptr, x float32, y float32) Result
var tvg_paint_set_transform func(paint uintptr, m *Matrix) Result
var tvg_paint_get_transform func(paint uintptr, m *Matrix) Result
var tvg_paint_set_opacity func(paint uintptr, opacity uint8) Result
var tvg_paint_get_opacity func(paint uintptr, opacity *uint8) Result
var tvg_paint_duplicate func(paint uintptr) uintptr
var tvg_paint_intersects func(paint uintptr, x int32, y int32, w int32, h int32) bool
var tvg_paint_get_aabb func(paint uintptr, x *float32, y *float32, w *float32, h *float32) Result
var tvg_paint_get_obb func(paint uintptr, pt4 *Point) Result
var tvg_paint_set_mask_method func(paint uintptr, target uintptr, method MaskMethod) Result
var tvg_paint_get_mask_method func(paint uintptr, target uintptr, method *MaskMethod) Result
var tvg_paint_set_clip func(paint uintptr, clipper uintptr) Result
var tvg_paint_get_clip func(paint uintptr) uintptr
var tvg_paint_get_parent func(paint uintptr) uintptr
var tvg_paint_get_type func(paint uintptr, typ *Type) Result
var tvg_paint_set_blend_method func(paint uintptr, method BlendMethod) Result

// Shape API

var tvg_shape_new func() uintptr
var tvg_shape_reset func(paint uintptr) Result
var tvg_shape_move_to func(paint uintptr, x float32, y float32) Result
var tvg_shape_line_to func(paint uintptr, x float32, y float32) Result
var tvg_shape_cubic_to func(paint uintptr, cx1 float32, cy1 float32, cx2 float32, cy2 float32, x float32, y float32) Result
var tvg_shape_close func(paint uintptr) Result
var tvg_shape_append_rect func(paint uintptr, x float32, y float32, w float32, h float32, rx float32, ry float32, cw bool) Result
var tvg_shape_append_circle func(paint uintptr, cx float32, cy float32, rx float32, ry float32, cw bool) Result
var tvg_shape_append_path func(paint uintptr, cmds *PathCommand, cmdCnt uint32, pts *Point, ptsCnt uint32) Result
var tvg_shape_get_path func(paint uintptr, cmds **PathCommand, cmdsCnt *uint32, pts **Point, ptsCnt *uint32) Result
var tvg_shape_set_stroke_width func(paint uintptr, width float32) Result
var tvg_shape_get_stroke_width func(paint uintptr, width *float32) Result
var tvg_shape_set_stroke_color func(paint uintptr, r uint8, g uint8, b uint8, a uint8) Result
var tvg_shape_get_stroke_color func(paint uintptr, r *uint8, g *uint8, b *uint8, a *uint8) Result
var tvg_shape_set_stroke_gradient func(paint uintptr, grad uintptr) Result
var tvg_shape_get_stroke_gradient func(paint uintptr, grad *uintptr) Result
var tvg_shape_set_stroke_dash func(paint uintptr, dashPattern *float32, cnt uint32, offset float32) Result
var tvg_shape_get_stroke_dash func(paint uintptr, dashPattern **float32, cnt *uint32, offset *float32) Result
var tvg_shape_set_stroke_cap func(paint uintptr, strokeCap StrokeCap) Result
var tvg_shape_get_stroke_cap func(paint uintptr, strokeCap *StrokeCap) Result
var tvg_shape_set_stroke_join func(paint uintptr, join StrokeJoin) Result
var tvg_shape_get_stroke_join func(paint uintptr, join *StrokeJoin) Result
var tvg_shape_set_stroke_miterlimit func(paint uintptr, miterlimit float32) Result
var tvg_shape_get_stroke_miterlimit func(paint uintptr, miterlimit *float32) Result
var tvg_shape_set_trimpath func(paint uintptr, begin float32, end float32, simultaneous bool) Result
var tvg_shape_set_fill_color func(paint uintptr, r uint8, g uint8, b uint8, a uint8) Result
var tvg_shape_get_fill_color func(paint uintptr, r *uint8, g *uint8, b *uint8, a *uint8) Result
var tvg_shape_set_fill_rule func(paint uintptr, rule FillRule) Result
var tvg_shape_get_fill_rule func(paint uintptr, rule *FillRule) Result
var tvg_shape_set_paint_order func(paint uintptr, strokeFirst bool) Result
var tvg_shape_set_gradient func(paint uintptr, grad uintptr) Result
var tvg_shape_get_gradient func(paint uintptr, grad *uintptr) Result

// Gradient API
var tvg_radial_gradient_new func() uintptr
var tvg_linear_gradient_new func() uintptr
var tvg_linear_gradient_set func(grad uintptr, x1 float32, y1 float32, x2 float32, y2 float32) Result
var tvg_linear_gradient_get func(grad uintptr, x1 *float32, y1 *float32, x2 *float32, y2 *float32) Result
var tvg_radial_gradient_set func(grad uintptr, cx float32, cy float32, r float32, fx float32, fy float32, fr float32) Result
var tvg_radial_gradient_get func(grad uintptr, cx *float32, cy *float32, r *float32, fx *float32, fy *float32, fr *float32) Result
var tvg_gradient_set_color_stops func(grad uintptr, color_stop *ColorStop, cnt uint32) Result
var tvg_gradient_get_color_stops func(grad uintptr, color_stop **ColorStop, cnt *uint32) Result
var tvg_gradient_set_spread func(grad uintptr, spread StrokeFill) Result
var tvg_gradient_get_spread func(grad uintptr, spread *StrokeFill) Result
var tvg_gradient_set_transform func(grad uintptr, m *Matrix) Result
var tvg_gradient_get_transform func(grad uintptr, m *Matrix) Result
var tvg_gradient_get_type func(grad uintptr, typ *Type) Result
var tvg_gradient_duplicate func(grad uintptr) uintptr
var tvg_gradient_del func(grad uintptr) Result

// Picture API
var tvg_picture_new func() uintptr
var tvg_picture_load func(picture uintptr, path string) Result
var tvg_picture_load_raw func(picture uintptr, data *uint32, w uint32, h uint32, cs ColorSpace, copy_ bool) Result
var tvg_picture_load_data func(picture uintptr, data *byte, size uint32, mimetype string, rpath string, copy_ bool) Result
var tvg_picture_set_asset_resolver func(picture uintptr, resolver uintptr, data uintptr) Result
var tvg_picture_set_size func(picture uintptr, w float32, h float32) Result
var tvg_picture_get_size func(picture uintptr, w *float32, h *float32) Result
var tvg_picture_set_origin func(picture uintptr, x float32, y float32) Result
var tvg_picture_get_origin func(picture uintptr, x *float32, y *float32) Result
var tvg_picture_get_paint func(picture uintptr, id uint32) uintptr
var tvg_picture_set_filter func(paint uintptr, method FilterMethod) Result
var tvg_picture_set_accessible func(paint uintptr, accessible bool) Result

// Scene API
var tvg_scene_new func() uintptr
var tvg_scene_add func(scene uintptr, paint uintptr) Result
var tvg_scene_insert func(scene uintptr, target uintptr, at uintptr) Result
var tvg_scene_remove func(scene uintptr, paint uintptr) Result
var tvg_scene_clear_effects func(scene uintptr) Result
var tvg_scene_add_effect_gaussian_blur func(scene uintptr, sigma float64, direction int32, border int32, quality int32) Result
var tvg_scene_add_effect_drop_shadow func(scene uintptr, r int32, g int32, b int32, a int32, angle float64, distance float64, sigma float64, quality int32) Result
var tvg_scene_add_effect_fill func(scene uintptr, r int32, g int32, b int32, a int32) Result
var tvg_scene_add_effect_tint func(scene uintptr, black_r int32, black_g int32, black_b int32, white_r int32, white_g int32, white_b int32, intensity float64) Result
var tvg_scene_add_effect_tritone func(scene uintptr, shadow_r int32, shadow_g int32, shadow_b int32, midtone_r int32, midtone_g int32, midtone_b int32, highlight_r int32, highlight_g int32, highlight_b int32, blend int32) Result

// Text API
var tvg_text_new func() uintptr
var tvg_text_set_font func(text uintptr, name string) Result
var tvg_text_set_size func(text uintptr, size float32) Result
var tvg_text_set_text func(text uintptr, utf8 string) Result
var tvg_text_align func(text uintptr, x float32, y float32) Result
var tvg_text_layout func(text uintptr, w float32, h float32) Result
var tvg_text_wrap_mode func(text uintptr, mode TextWrap) Result
var tvg_text_spacing func(text uintptr, letter float32, line float32) Result
var tvg_text_set_italic func(text uintptr, shear float32) Result
var tvg_text_set_outline func(text uintptr, width float32, r uint8, g uint8, b uint8) Result
var tvg_text_set_color func(text uintptr, r uint8, g uint8, b uint8) Result
var tvg_text_set_gradient func(text uintptr, gradient uintptr) Result
var tvg_text_get_text_metrics func(text uintptr, metrics *TextMetrics) Result
var tvg_text_get_glyph_metrics func(text uintptr, ch *byte, metrics *GlyphMetrics) Result
var tvg_font_load func(path string) Result
var tvg_font_load_data func(name string, data *byte, size uint32, mimetype string, copy_ bool) Result
var tvg_font_unload func(path string) Result

// Saver API
var tvg_saver_new func() Saver
var tvg_saver_save_paint func(saver uintptr, paint uintptr, path string, quality uint32) Result
var tvg_saver_save_animation func(saver uintptr, animation uintptr, path string, quality uint32, fps uint32) Result
var tvg_saver_sync func(saver uintptr) Result
var tvg_saver_del func(saver uintptr) Result

// Animation API
var tvg_animation_new func() uintptr
var tvg_animation_set_frame func(animation Animation, no float32) Result
var tvg_animation_get_picture func(animation Animation) uintptr
var tvg_animation_get_frame func(animation Animation, no *float32) Result
var tvg_animation_get_total_frame func(animation Animation, cnt *float32) Result
var tvg_animation_get_duration func(animation Animation, duration *float32) Result
var tvg_animation_set_segment func(animation Animation, begin float32, end float32) Result
var tvg_animation_get_segment func(animation Animation, begin *float32, end *float32) Result
var tvg_animation_del func(animation Animation) Result

// Accessor API
var tvg_accessor_new func() uintptr
var tvg_accessor_del func(accessor uintptr) Result
var tvg_accessor_set func(accessor uintptr, paint uintptr, func_ uintptr, data uintptr) Result
var tvg_accessor_generate_id func(name string) uint32
var tvg_accessor_get_name func(accessor uintptr, id uint32) string

// LottieAnimation API
var tvg_lottie_animation_new func() uintptr
var tvg_lottie_animation_gen_slot func(animation uintptr, slot string) uint32
var tvg_lottie_animation_apply_slot func(animation uintptr, id uint32) Result
var tvg_lottie_animation_del_slot func(animation uintptr, id uint32) Result
var tvg_lottie_animation_set_marker func(animation uintptr, marker string) Result
var tvg_lottie_animation_get_markers_cnt func(animation uintptr, cnt *uint32) Result
var tvg_lottie_animation_get_marker func(animation uintptr, idx uint32, name **byte) Result
var tvg_lottie_animation_get_marker_info func(animation uintptr, idx uint32, name **byte, begin *float32, end *float32) Result
var tvg_lottie_animation_tween func(animation uintptr, from float32, to float32, progress float32) Result
var tvg_lottie_animation_set_quality func(animation uintptr, value uint8) Result
var tvg_lottie_animation_set_audio_resolver func(animation uintptr, resolver uintptr, data uintptr) Result

var libThorvgInitialised = false

func initLibThorvg() error {
	if libThorvgInitialised {
		return nil
	}
	defer func() { libThorvgInitialised = true }()

	filepath, err := libraryFilepath()
	if err != nil {
		return err
	}

	lib, err := purego.Dlopen(filepath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return err
	}

	// Engine API
	purego.RegisterLibFunc(&tvg_engine_init, lib, "tvg_engine_init")
	purego.RegisterLibFunc(&tvg_engine_term, lib, "tvg_engine_term")
	purego.RegisterLibFunc(&tvg_engine_version, lib, "tvg_engine_version")

	// Canvas API
	purego.RegisterLibFunc(&tvg_swcanvas_create, lib, "tvg_swcanvas_create")
	purego.RegisterLibFunc(&tvg_swcanvas_set_target, lib, "tvg_swcanvas_set_target")
	purego.RegisterLibFunc(&tvg_glcanvas_create, lib, "tvg_glcanvas_create")
	purego.RegisterLibFunc(&tvg_glcanvas_set_target, lib, "tvg_glcanvas_set_target")
	purego.RegisterLibFunc(&tvg_wgcanvas_create, lib, "tvg_wgcanvas_create")
	purego.RegisterLibFunc(&tvg_wgcanvas_set_target, lib, "tvg_wgcanvas_set_target")
	purego.RegisterLibFunc(&tvg_canvas_destroy, lib, "tvg_canvas_destroy")
	purego.RegisterLibFunc(&tvg_canvas_add, lib, "tvg_canvas_add")
	purego.RegisterLibFunc(&tvg_canvas_insert, lib, "tvg_canvas_insert")
	purego.RegisterLibFunc(&tvg_canvas_remove, lib, "tvg_canvas_remove")
	purego.RegisterLibFunc(&tvg_canvas_update, lib, "tvg_canvas_update")
	purego.RegisterLibFunc(&tvg_canvas_draw, lib, "tvg_canvas_draw")
	purego.RegisterLibFunc(&tvg_canvas_sync, lib, "tvg_canvas_sync")
	purego.RegisterLibFunc(&tvg_canvas_set_viewport, lib, "tvg_canvas_set_viewport")

	// Paint API
	purego.RegisterLibFunc(&tvg_paint_rel, lib, "tvg_paint_rel")
	purego.RegisterLibFunc(&tvg_paint_ref, lib, "tvg_paint_ref")
	purego.RegisterLibFunc(&tvg_paint_unref, lib, "tvg_paint_unref")
	purego.RegisterLibFunc(&tvg_paint_get_ref, lib, "tvg_paint_get_ref")
	purego.RegisterLibFunc(&tvg_paint_set_visible, lib, "tvg_paint_set_visible")
	purego.RegisterLibFunc(&tvg_paint_get_visible, lib, "tvg_paint_get_visible")
	purego.RegisterLibFunc(&tvg_paint_get_id, lib, "tvg_paint_get_id")
	purego.RegisterLibFunc(&tvg_paint_set_id, lib, "tvg_paint_set_id")
	purego.RegisterLibFunc(&tvg_paint_scale, lib, "tvg_paint_scale")
	purego.RegisterLibFunc(&tvg_paint_rotate, lib, "tvg_paint_rotate")
	purego.RegisterLibFunc(&tvg_paint_translate, lib, "tvg_paint_translate")
	purego.RegisterLibFunc(&tvg_paint_set_transform, lib, "tvg_paint_set_transform")
	purego.RegisterLibFunc(&tvg_paint_get_transform, lib, "tvg_paint_get_transform")
	purego.RegisterLibFunc(&tvg_paint_set_opacity, lib, "tvg_paint_set_opacity")
	purego.RegisterLibFunc(&tvg_paint_get_opacity, lib, "tvg_paint_get_opacity")
	purego.RegisterLibFunc(&tvg_paint_duplicate, lib, "tvg_paint_duplicate")
	purego.RegisterLibFunc(&tvg_paint_intersects, lib, "tvg_paint_intersects")
	purego.RegisterLibFunc(&tvg_paint_get_aabb, lib, "tvg_paint_get_aabb")
	purego.RegisterLibFunc(&tvg_paint_get_obb, lib, "tvg_paint_get_obb")
	purego.RegisterLibFunc(&tvg_paint_set_mask_method, lib, "tvg_paint_set_mask_method")
	purego.RegisterLibFunc(&tvg_paint_get_mask_method, lib, "tvg_paint_get_mask_method")
	purego.RegisterLibFunc(&tvg_paint_set_clip, lib, "tvg_paint_set_clip")
	purego.RegisterLibFunc(&tvg_paint_get_clip, lib, "tvg_paint_get_clip")
	purego.RegisterLibFunc(&tvg_paint_get_parent, lib, "tvg_paint_get_parent")
	purego.RegisterLibFunc(&tvg_paint_get_type, lib, "tvg_paint_get_type")
	purego.RegisterLibFunc(&tvg_paint_set_blend_method, lib, "tvg_paint_set_blend_method")

	// Shape API
	purego.RegisterLibFunc(&tvg_shape_new, lib, "tvg_shape_new")
	purego.RegisterLibFunc(&tvg_shape_reset, lib, "tvg_shape_reset")
	purego.RegisterLibFunc(&tvg_shape_move_to, lib, "tvg_shape_move_to")
	purego.RegisterLibFunc(&tvg_shape_line_to, lib, "tvg_shape_line_to")
	purego.RegisterLibFunc(&tvg_shape_cubic_to, lib, "tvg_shape_cubic_to")
	purego.RegisterLibFunc(&tvg_shape_close, lib, "tvg_shape_close")
	purego.RegisterLibFunc(&tvg_shape_append_rect, lib, "tvg_shape_append_rect")
	purego.RegisterLibFunc(&tvg_shape_append_circle, lib, "tvg_shape_append_circle")
	purego.RegisterLibFunc(&tvg_shape_append_path, lib, "tvg_shape_append_path")
	purego.RegisterLibFunc(&tvg_shape_get_path, lib, "tvg_shape_get_path")
	purego.RegisterLibFunc(&tvg_shape_set_stroke_width, lib, "tvg_shape_set_stroke_width")
	purego.RegisterLibFunc(&tvg_shape_get_stroke_width, lib, "tvg_shape_get_stroke_width")
	purego.RegisterLibFunc(&tvg_shape_set_stroke_color, lib, "tvg_shape_set_stroke_color")
	purego.RegisterLibFunc(&tvg_shape_get_stroke_color, lib, "tvg_shape_get_stroke_color")
	purego.RegisterLibFunc(&tvg_shape_set_stroke_gradient, lib, "tvg_shape_set_stroke_gradient")
	purego.RegisterLibFunc(&tvg_shape_get_stroke_gradient, lib, "tvg_shape_get_stroke_gradient")
	purego.RegisterLibFunc(&tvg_shape_set_stroke_dash, lib, "tvg_shape_set_stroke_dash")
	purego.RegisterLibFunc(&tvg_shape_get_stroke_dash, lib, "tvg_shape_get_stroke_dash")
	purego.RegisterLibFunc(&tvg_shape_set_stroke_cap, lib, "tvg_shape_set_stroke_cap")
	purego.RegisterLibFunc(&tvg_shape_get_stroke_cap, lib, "tvg_shape_get_stroke_cap")
	purego.RegisterLibFunc(&tvg_shape_set_stroke_join, lib, "tvg_shape_set_stroke_join")
	purego.RegisterLibFunc(&tvg_shape_get_stroke_join, lib, "tvg_shape_get_stroke_join")
	purego.RegisterLibFunc(&tvg_shape_set_stroke_miterlimit, lib, "tvg_shape_set_stroke_miterlimit")
	purego.RegisterLibFunc(&tvg_shape_get_stroke_miterlimit, lib, "tvg_shape_get_stroke_miterlimit")
	purego.RegisterLibFunc(&tvg_shape_set_trimpath, lib, "tvg_shape_set_trimpath")
	purego.RegisterLibFunc(&tvg_shape_set_fill_color, lib, "tvg_shape_set_fill_color")
	purego.RegisterLibFunc(&tvg_shape_get_fill_color, lib, "tvg_shape_get_fill_color")
	purego.RegisterLibFunc(&tvg_shape_set_fill_rule, lib, "tvg_shape_set_fill_rule")
	purego.RegisterLibFunc(&tvg_shape_get_fill_rule, lib, "tvg_shape_get_fill_rule")
	purego.RegisterLibFunc(&tvg_shape_set_paint_order, lib, "tvg_shape_set_paint_order")
	purego.RegisterLibFunc(&tvg_shape_set_gradient, lib, "tvg_shape_set_gradient")
	purego.RegisterLibFunc(&tvg_shape_get_gradient, lib, "tvg_shape_get_gradient")

	// Gradient API
	purego.RegisterLibFunc(&tvg_radial_gradient_new, lib, "tvg_radial_gradient_new")
	purego.RegisterLibFunc(&tvg_linear_gradient_new, lib, "tvg_linear_gradient_new")
	purego.RegisterLibFunc(&tvg_linear_gradient_set, lib, "tvg_linear_gradient_set")
	purego.RegisterLibFunc(&tvg_linear_gradient_get, lib, "tvg_linear_gradient_get")
	purego.RegisterLibFunc(&tvg_radial_gradient_set, lib, "tvg_radial_gradient_set")
	purego.RegisterLibFunc(&tvg_radial_gradient_get, lib, "tvg_radial_gradient_get")
	purego.RegisterLibFunc(&tvg_gradient_set_color_stops, lib, "tvg_gradient_set_color_stops")
	purego.RegisterLibFunc(&tvg_gradient_get_color_stops, lib, "tvg_gradient_get_color_stops")
	purego.RegisterLibFunc(&tvg_gradient_set_spread, lib, "tvg_gradient_set_spread")
	purego.RegisterLibFunc(&tvg_gradient_get_spread, lib, "tvg_gradient_get_spread")
	purego.RegisterLibFunc(&tvg_gradient_set_transform, lib, "tvg_gradient_set_transform")
	purego.RegisterLibFunc(&tvg_gradient_get_transform, lib, "tvg_gradient_get_transform")
	purego.RegisterLibFunc(&tvg_gradient_get_type, lib, "tvg_gradient_get_type")
	purego.RegisterLibFunc(&tvg_gradient_duplicate, lib, "tvg_gradient_duplicate")
	purego.RegisterLibFunc(&tvg_gradient_del, lib, "tvg_gradient_del")

	// Picture API
	purego.RegisterLibFunc(&tvg_picture_new, lib, "tvg_picture_new")
	purego.RegisterLibFunc(&tvg_picture_load, lib, "tvg_picture_load")
	purego.RegisterLibFunc(&tvg_picture_load_raw, lib, "tvg_picture_load_raw")
	purego.RegisterLibFunc(&tvg_picture_load_data, lib, "tvg_picture_load_data")
	purego.RegisterLibFunc(&tvg_picture_set_asset_resolver, lib, "tvg_picture_set_asset_resolver")
	purego.RegisterLibFunc(&tvg_picture_set_size, lib, "tvg_picture_set_size")
	purego.RegisterLibFunc(&tvg_picture_get_size, lib, "tvg_picture_get_size")
	purego.RegisterLibFunc(&tvg_picture_set_origin, lib, "tvg_picture_set_origin")
	purego.RegisterLibFunc(&tvg_picture_get_origin, lib, "tvg_picture_get_origin")
	purego.RegisterLibFunc(&tvg_picture_get_paint, lib, "tvg_picture_get_paint")
	purego.RegisterLibFunc(&tvg_picture_set_filter, lib, "tvg_picture_set_filter")
	purego.RegisterLibFunc(&tvg_picture_set_accessible, lib, "tvg_picture_set_accessible")

	// Scene API
	purego.RegisterLibFunc(&tvg_scene_new, lib, "tvg_scene_new")
	purego.RegisterLibFunc(&tvg_scene_add, lib, "tvg_scene_add")
	purego.RegisterLibFunc(&tvg_scene_insert, lib, "tvg_scene_insert")
	purego.RegisterLibFunc(&tvg_scene_remove, lib, "tvg_scene_remove")
	purego.RegisterLibFunc(&tvg_scene_clear_effects, lib, "tvg_scene_clear_effects")
	purego.RegisterLibFunc(&tvg_scene_add_effect_gaussian_blur, lib, "tvg_scene_add_effect_gaussian_blur")
	purego.RegisterLibFunc(&tvg_scene_add_effect_drop_shadow, lib, "tvg_scene_add_effect_drop_shadow")
	purego.RegisterLibFunc(&tvg_scene_add_effect_fill, lib, "tvg_scene_add_effect_fill")
	purego.RegisterLibFunc(&tvg_scene_add_effect_tint, lib, "tvg_scene_add_effect_tint")
	purego.RegisterLibFunc(&tvg_scene_add_effect_tritone, lib, "tvg_scene_add_effect_tritone")

	// Text API
	purego.RegisterLibFunc(&tvg_text_new, lib, "tvg_text_new")
	purego.RegisterLibFunc(&tvg_text_set_font, lib, "tvg_text_set_font")
	purego.RegisterLibFunc(&tvg_text_set_size, lib, "tvg_text_set_size")
	purego.RegisterLibFunc(&tvg_text_set_text, lib, "tvg_text_set_text")
	purego.RegisterLibFunc(&tvg_text_align, lib, "tvg_text_align")
	purego.RegisterLibFunc(&tvg_text_layout, lib, "tvg_text_layout")
	purego.RegisterLibFunc(&tvg_text_wrap_mode, lib, "tvg_text_wrap_mode")
	purego.RegisterLibFunc(&tvg_text_spacing, lib, "tvg_text_spacing")
	purego.RegisterLibFunc(&tvg_text_set_italic, lib, "tvg_text_set_italic")
	purego.RegisterLibFunc(&tvg_text_set_outline, lib, "tvg_text_set_outline")
	purego.RegisterLibFunc(&tvg_text_set_color, lib, "tvg_text_set_color")
	// tvg_text_set_gradient is dependent on the linux struct argument support
	// in https://github.com/ebitengine/purego/pull/361.
	purego.RegisterLibFunc(&tvg_text_set_gradient, lib, "tvg_text_set_gradient")
	purego.RegisterLibFunc(&tvg_text_get_text_metrics, lib, "tvg_text_get_text_metrics")
	purego.RegisterLibFunc(&tvg_text_get_glyph_metrics, lib, "tvg_text_get_glyph_metrics")
	purego.RegisterLibFunc(&tvg_font_load, lib, "tvg_font_load")
	purego.RegisterLibFunc(&tvg_font_load_data, lib, "tvg_font_load_data")
	purego.RegisterLibFunc(&tvg_font_unload, lib, "tvg_font_unload")

	// Saver API
	purego.RegisterLibFunc(&tvg_saver_new, lib, "tvg_saver_new")
	purego.RegisterLibFunc(&tvg_saver_save_paint, lib, "tvg_saver_save_paint")
	purego.RegisterLibFunc(&tvg_saver_save_animation, lib, "tvg_saver_save_animation")
	purego.RegisterLibFunc(&tvg_saver_sync, lib, "tvg_saver_sync")
	purego.RegisterLibFunc(&tvg_saver_del, lib, "tvg_saver_del")

	// Animation API
	purego.RegisterLibFunc(&tvg_animation_new, lib, "tvg_animation_new")
	purego.RegisterLibFunc(&tvg_animation_set_frame, lib, "tvg_animation_set_frame")
	purego.RegisterLibFunc(&tvg_animation_get_picture, lib, "tvg_animation_get_picture")
	purego.RegisterLibFunc(&tvg_animation_get_frame, lib, "tvg_animation_get_frame")
	purego.RegisterLibFunc(&tvg_animation_get_total_frame, lib, "tvg_animation_get_total_frame")
	purego.RegisterLibFunc(&tvg_animation_get_duration, lib, "tvg_animation_get_duration")
	purego.RegisterLibFunc(&tvg_animation_set_segment, lib, "tvg_animation_set_segment")
	purego.RegisterLibFunc(&tvg_animation_get_segment, lib, "tvg_animation_get_segment")
	purego.RegisterLibFunc(&tvg_animation_del, lib, "tvg_animation_del")

	// Accessor API
	purego.RegisterLibFunc(&tvg_accessor_new, lib, "tvg_accessor_new")
	purego.RegisterLibFunc(&tvg_accessor_del, lib, "tvg_accessor_del")
	purego.RegisterLibFunc(&tvg_accessor_set, lib, "tvg_accessor_set")
	purego.RegisterLibFunc(&tvg_accessor_generate_id, lib, "tvg_accessor_generate_id")
	purego.RegisterLibFunc(&tvg_accessor_get_name, lib, "tvg_accessor_get_name")

	// LottieAnimation API
	purego.RegisterLibFunc(&tvg_lottie_animation_new, lib, "tvg_lottie_animation_new")
	purego.RegisterLibFunc(&tvg_lottie_animation_gen_slot, lib, "tvg_lottie_animation_gen_slot")
	purego.RegisterLibFunc(&tvg_lottie_animation_apply_slot, lib, "tvg_lottie_animation_apply_slot")
	purego.RegisterLibFunc(&tvg_lottie_animation_del_slot, lib, "tvg_lottie_animation_del_slot")
	purego.RegisterLibFunc(&tvg_lottie_animation_set_marker, lib, "tvg_lottie_animation_set_marker")
	purego.RegisterLibFunc(&tvg_lottie_animation_get_markers_cnt, lib, "tvg_lottie_animation_get_markers_cnt")
	purego.RegisterLibFunc(&tvg_lottie_animation_get_marker, lib, "tvg_lottie_animation_get_marker")
	purego.RegisterLibFunc(&tvg_lottie_animation_get_marker_info, lib, "tvg_lottie_animation_get_marker_info")
	purego.RegisterLibFunc(&tvg_lottie_animation_tween, lib, "tvg_lottie_animation_tween")
	purego.RegisterLibFunc(&tvg_lottie_animation_set_quality, lib, "tvg_lottie_animation_set_quality")
	purego.RegisterLibFunc(&tvg_lottie_animation_set_audio_resolver, lib, "tvg_lottie_animation_set_audio_resolver")

	return nil
}
