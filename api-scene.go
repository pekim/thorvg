package thorvg

type Scene struct {
	paintCommon
}

// Creates a new Scene object.
//
// This function allocates and returns a new Scene instance. To properly destroy the Scene object, use [Paint.Rel]
//
// Return: The newly created Scene object.
//
// see [Paint.Rel]
func SceneNew() Scene {
	return Scene{
		paintCommon: paintCommon{
			paint_: tvg_scene_new(),
		},
	}
}

// Adds a paint object to the scene.
//
// Appends the specified paint object to the given scene. Only paint objects added to the scene are considered rendering targets.
//
//   - scene : A handle to the scene object.                  This parameter must not be nullptr.
//   - paint : A handle to the paint object to be added to the scene.                  This parameter must not be nullptr.
//
// Note: Ownership of the paint object is transferred to the canvas upon       successful addition. To retain ownership, call [Paint.Ref]       before adding it to the scene.
//
// Note: The rendering order of paint objects follows their order in the root       scene. If layering is required, ensure the paints are added in the       desired order.
//
// see [Scene.Insert]
//
// see [Scene.Remove]
//
// Since: 1.0
func (scene Scene) Add(paint Paint) error {
	return tvg_scene_add(scene.paint_, paint.paint()).error()
}

// Inserts a paint object into the scene.
//
// Inserts the specified paint object into the scene immediately before the given paint object at. The at parameter must reference an existing paint object already added to the scene.
//
//   - scene :  A handle to the scene object.                   This parameter must not be nullptr.
//   - target : A handle to the paint object to be inserted into the scene.                   This parameter must not be nullptr.
//   - at :     A handle to an existing paint object in the scene before                   which target will be inserted.                   This parameter must not be nullptr.
//
// Note: Ownership of the target object is transferred to the scene upon       successful addition. To retain ownership, call [Paint.Ref]       before adding it to the scene.
//
// Note: The rendering order of paint objects follows their order in the root       scene. If layering is required, ensure the paints are added in the       desired order.
//
// see [Scene.Add]
//
// see [Scene.Remove]
//
// Since: 1.0
func (scene Scene) Insert(target Paint, at Paint) error {
	return tvg_scene_insert(scene.paint_, target.paint(), at.paint()).error()
}

// Removes a paint object from the scene.
//
// This function removes a specified paint object from the scene. If no paint object is specified (i.e., the default nullptr is used), the function performs to clear all paints from the scene.
//
//   - scene : The scene object.
//   - paint : The paint object to be removed from the scene.                  If nullptr, remove all the paints from the scene.
//
// see [Scene.Add]
// Since: 1.0
func (scene Scene) Remove(paint Paint) error {
	return tvg_scene_remove(scene.paint_, paint.paint()).error()
}

// Clears all previously applied scene effects.
//
// This function clears all effects that have been applied to the scene, restoring it to its original state without any post-processing.
//
//   - scene : The scene object.
//
// Since: 1.0
func (scene Scene) ClearEffects() error {
	return tvg_scene_clear_effects(scene.paint_).error()
}

// Adds a Gaussian blur effect to the scene.
//
// This function adds a Gaussian blur filter to the scene as a post-processing effect. The blur can be applied in different directions with configurable border handling and quality settings.
//
//   - scene : The scene object.
//   - sigma : The blur radius (sigma) value. Must be >= 0.
//   - direction : Blur direction: 0 = both directions, 1 = horizontal only, 2 = vertical only.
//   - border : Reserved for border handling. Currently ignored.
//   - quality : Visual quality level [0 - 100].
//
// Since: 1.0
func (scene Scene) AddEffectGaussianBlur(sigma float64, direction int, border int, quality int) error {
	return tvg_scene_add_effect_gaussian_blur(
		scene.paint_, sigma,
		int32(direction), int32(border), int32(quality),
	).error()
}

// Adds a centered linear motion blur.
//
//   - scene : The scene to apply the effect to.
//   - distance : Total blur length in scene-local units. Must be >= 0.
//   - angle : Angle in degrees, clockwise from the positive X axis. The scene transform applies to the blur vector.
//   - quality : Visual quality level [0 - 100].
//
// Note: Experimental API
func (scene Scene) AddEffectMotionBlur(distance float64, angle float64, quality int32) error {
	return tvg_scene_add_effect_motion_blur(scene.paint_, distance, angle, quality).error()
}

// Adds a drop shadow effect to the scene.
//
// This function adds a drop shadow with a Gaussian blur to the scene. The shadow can be customized using color, opacity, angle, distance, blur radius (sigma), and quality parameters.
//
//   - scene : The scene object.
//   - r : Red channel value of the shadow color [0 - 255].
//   - g : Green channel value of the shadow color [0 - 255].
//   - b : Blue channel value of the shadow color [0 - 255].
//   - a : Alpha (opacity) channel value of the shadow [0 - 255].
//   - angle : Shadow direction in degrees [0 - 360].
//   - distance : Distance of the shadow from the original object.
//   - sigma : Gaussian blur sigma value for the shadow. Must be > 0.
//   - quality : Visual quality level [0 - 100].
//
// Since: 1.0
func (scene Scene) AddEffectDropShadow(
	r int, g int, b int, a int,
	angle float64, distance float64, sigma float64, quality int,
) error {
	return tvg_scene_add_effect_drop_shadow(scene.paint_,
		int32(r), int32(g), int32(b), int32(a),
		angle, distance, sigma, int32(quality),
	).error()
}

// Adds a fill color effect to the scene.
//
// This function overrides the scene's content colors with the specified fill color.
//
//   - scene : The scene object.
//   - r : Red color channel value [0 - 255].
//   - g : Green color channel value [0 - 255].
//   - b : Blue color channel value [0 - 255].
//   - a : Alpha (opacity) channel value [0 - 255].
//
// Since: 1.0
func (scene Scene) AddEffectFill(r int, g int, b int, a int) error {
	return tvg_scene_add_effect_fill(scene.paint_, int32(r), int32(g), int32(b), int32(a)).error()
}

// Adds a tint effect to the scene.
//
// This function tints the current scene using specified black and white color values, modulated by a given intensity.
//
//   - scene : The scene object.
//   - black_r : Red component of the black color [0 - 255].
//   - black_g : Green component of the black color [0 - 255].
//   - black_b : Blue component of the black color [0 - 255].
//   - white_r : Red component of the white color [0 - 255].
//   - white_g : Green component of the white color [0 - 255].
//   - white_b : Blue component of the white color [0 - 255].
//   - intensity : Tint intensity value [0 - 100].
//
// Since: 1.0
func (scene Scene) AddEffectTint(
	black_r int, black_g int, black_b int,
	white_r int, white_g int, white_b int,
	intensity float64,
) error {
	return tvg_scene_add_effect_tint(scene.paint_,
		int32(black_r), int32(black_g), int32(black_b),
		int32(white_r), int32(white_g), int32(white_b),
		intensity,
	).error()
}

// Adds a tritone color effect to the scene.
//
// This function adds a tritone color effect to the given scene using three sets of RGB values  representing shadow, midtone, and highlight colors.
//
//   - scene : The scene object.
//   - shadow_r : Red component of the shadow color [0 - 255].
//   - shadow_g : Green component of the shadow color [0 - 255].
//   - shadow_b : Blue component of the shadow color [0 - 255].
//   - midtone_r : Red component of the midtone color [0 - 255].
//   - midtone_g : Green component of the midtone color [0 - 255].
//   - midtone_b : Blue component of the midtone color [0 - 255].
//   - highlight_r : Red component of the highlight color [0 - 255].
//   - highlight_g : Green component of the highlight color [0 - 255].
//   - highlight_b : Blue component of the highlight color [0 - 255].
//   - blend : A blending factor that determines the mix between the original color and the tritone colors [0 - 255].
//
// Since: 1.0
func (scene Scene) AddEffectTritone(
	shadow_r int, shadow_g int, shadow_b int,
	midtone_r int, midtone_g int, midtone_b int,
	highlight_r int, highlight_g int, highlight_b int,
	blend int,
) error {
	return tvg_scene_add_effect_tritone(scene.paint_,
		int32(shadow_r), int32(shadow_g), int32(shadow_b),
		int32(midtone_r), int32(midtone_g), int32(midtone_b),
		int32(highlight_r), int32(highlight_g), int32(highlight_b),
		int32(blend),
	).error()
}

/*
Duplicate duplicates a Scene.

Creates a new object and sets its all properties as in the original object.

	@param[in] paint The Tvg_Paint object to be copied.

	@return A copied Tvg_Paint object if succeed, @c nullptr otherwise.
*/
func (scene Scene) Duplicate() Scene {
	return Scene{
		paintCommon: scene.duplicate(),
	}
}
